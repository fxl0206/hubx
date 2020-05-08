package envoy

import (
	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	pauth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	als "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	alf "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	lua "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/lua/v2"
	hcm "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	tcp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/types"
	ext "k8s.io/api/extensions/v1beta1"
	"strconv"

	"time"

	//xdsv1 "aif.io/hubx/k8s/portal/api/v1"
	"github.com/envoyproxy/go-control-plane/pkg/cache"

	"fmt"
	"strings"
)
const (

	// XdsCluster is the cluster name for the control server (used by non-ADS set-up)
	XdsCluster = "xds_cluster"

	// Ads mode for resources: one aggregated xDS service
	Ads = "ads"

	// Xds mode for resources: individual xDS services
	Xds = "xds"

	// Rest mode for resources: polling using Fetch
	Rest = "rest"
)

var (
	// RefreshDelay for the polling config source
	RefreshDelay = 500 * time.Millisecond
)

func MakeEndpoint(clusterName string,address string, port uint32) *v2.ClusterLoadAssignment {
	return &v2.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []endpoint.LocalityLbEndpoints{{
			LbEndpoints: []endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.TCP,
									Address:  address,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: port,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}


// MakeCluster creates a cluster using either ADS or EDS.
func MakeCluster(mode string, clusterName string) *v2.Cluster {
	var edsSource *core.ConfigSource
	switch mode {
	case Ads:
		edsSource = &core.ConfigSource{
			ConfigSourceSpecifier: &core.ConfigSource_Ads{
				Ads: &core.AggregatedConfigSource{},
			},
		}
	case Xds:
		edsSource = &core.ConfigSource{
			ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
				ApiConfigSource: &core.ApiConfigSource{
					ApiType: core.ApiConfigSource_GRPC,
					GrpcServices: []*core.GrpcService{{
						TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
							EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: XdsCluster},
						},
					}},
				},
			},
		}
	case Rest:
		edsSource = &core.ConfigSource{
			ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
				ApiConfigSource: &core.ApiConfigSource{
					ApiType:      core.ApiConfigSource_REST,
					ClusterNames: []string{XdsCluster},
					RefreshDelay: &RefreshDelay,
				},
			},
		}
	}

	return &v2.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       5 * time.Second,
		ClusterDiscoveryType: &v2.Cluster_Type{Type: v2.Cluster_EDS},
		EdsClusterConfig: &v2.Cluster_EdsClusterConfig{
			EdsConfig: edsSource,
		},
	}
}
func newRoute(clusterName,auth,prefix string,routes [] route.Route)  [] route.Route{
	if prefix == ""{
		prefix="/"
	}
	newRoute:=route.Route{
		Match: route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: prefix,
			},
		},
		Metadata:&core.Metadata{
			FilterMetadata:map[string]*types.Struct{
				"envoy.lua":&types.Struct{
					Fields: map[string]*types.Value{
						"credentials":&types.Value{
							Kind:&types.Value_ListValue{
								ListValue:&types.ListValue{
									Values:[]*types.Value{
										{
											Kind:&types.Value_StringValue{
												StringValue:"Basic "+auth+"=",
											},
										},
										{
											Kind:&types.Value_StringValue{
												StringValue:"Basic "+auth+"==",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: clusterName,
				},
			},
		},
	}
	return append(routes,newRoute)
}

func MakeVirtualHost(vsName ,domain string,routes [] route.Route) route.VirtualHost{
	return route.VirtualHost{
		Name:    vsName,
		Domains: []string{domain},
		Routes: routes,
	}
}
// MakeRoute creates an HTTP route that routes to a given cluster.
func MakeRoute(routeName string,virtualHosts [] route.VirtualHost) *v2.RouteConfiguration {
	return &v2.RouteConfiguration{
		Name: routeName,
		VirtualHosts: virtualHosts,
	}
}

// data source configuration
func configSource(mode string) *core.ConfigSource {
	source := &core.ConfigSource{}
	switch mode {
	case Ads:
		source.ConfigSourceSpecifier = &core.ConfigSource_Ads{
			Ads: &core.AggregatedConfigSource{},
		}
	case Xds:
		source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
			ApiConfigSource: &core.ApiConfigSource{
				ApiType: core.ApiConfigSource_GRPC,
				GrpcServices: []*core.GrpcService{{
					TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
						EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: XdsCluster},
					},
				}},
			},
		}
	case Rest:
		source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
			ApiConfigSource: &core.ApiConfigSource{
				ApiType:      core.ApiConfigSource_REST,
				ClusterNames: []string{XdsCluster},
				RefreshDelay: &RefreshDelay,
			},
		}
	}
	return source
}

// MakeHTTPListener creates a listener using either ADS or RDS for the route.
func MakeHTTPListener(ssl bool,auth,mode string, listenerName string, port uint32, route string) *v2.Listener {
	rdsSource := configSource(mode)

	// access log service configuration
	alsConfig := &als.HttpGrpcAccessLogConfig{
		CommonConfig: &als.CommonGrpcAccessLogConfig{
			LogName: "echo",
			GrpcService: &core.GrpcService{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
						ClusterName: XdsCluster,
					},
				},
			},
		},
	}
	alsConfigPbst, err := types.MarshalAny(alsConfig)
	if err != nil {
		panic(err)
	}
    httpFilters:=[]*hcm.HttpFilter{{
		Name: util.Router,
	}}

    if auth != "" {
    	zlua:=&lua.Lua{
    		InlineCode:`
                 function envoy_on_request(request_handle)
                    -- Surely you have to check if request_handle:metadata():get("credentials") has
                    -- nothing then you need to decide what to do.
                    for _, credential in pairs(request_handle:metadata():get("credentials")) do
                      if request_handle:headers():get("authorization") == credential
                      then
                        return
                      end
                    end
                    request_handle:respond(
                      {[":status"] = "401", ["WWW-Authenticate"] = "Basic realm=\"Unknown\""}, "Unauthorized"
                    )
                 end
             `,
		}
		llua, err := types.MarshalAny(zlua)
		if err != nil {
			panic(err)
		}
		httpFilters=[]*hcm.HttpFilter{
			{
			  Name:"envoy.lua",
			  ConfigType:&hcm.HttpFilter_TypedConfig{
			  	TypedConfig:llua,
			  },
		    },
			{
			  Name: util.Router,
		    },
		}
	}
	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.AUTO,
		StatPrefix: "http",
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    *rdsSource,
				RouteConfigName: route,
			},
		},
		HttpFilters: httpFilters,
		AccessLog: []*alf.AccessLog{{
			Name: util.HTTPGRPCAccessLog,
			ConfigType: &alf.AccessLog_TypedConfig{
				TypedConfig: alsConfigPbst,
			},
		}},
	}

	pbst, err := types.MarshalAny(manager)
	if err != nil {
		panic(err)
	}
    chain:=listener.FilterChain{
		Filters: []listener.Filter{{
			Name: util.HTTPConnectionManager,
			ConfigType: &listener.Filter_TypedConfig{
				TypedConfig: pbst,
			},
		}},
	}
    if ssl {
		chain.TlsContext=&pauth.DownstreamTlsContext{
			CommonTlsContext:&pauth.CommonTlsContext{
				TlsCertificates:[]*pauth.TlsCertificate{
					&pauth.TlsCertificate{
						CertificateChain:&core.DataSource{
							Specifier:&core.DataSource_Filename{
								"/etc/letsencrypt/fullchain1.pem",
							},
						},
						PrivateKey:&core.DataSource{
							Specifier:&core.DataSource_Filename{
								"/etc/letsencrypt/privkey1.pem",
							},
						},
					},
				},
			},
		}
	}
	return &v2.Listener{
		Name: listenerName,
		Address: core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: []listener.FilterChain{chain},
	}
}

// MakeTCPListener creates a TCP listener for a cluster.
func MakeTCPListener(listenerName string, port uint32, clusterName string) *v2.Listener {
	// TCP filter configuration
	config := &tcp.TcpProxy{
		StatPrefix: "tcp",
		ClusterSpecifier: &tcp.TcpProxy_Cluster{
			Cluster: clusterName,
		},
	}
	pbst, err := types.MarshalAny(config)
	if err != nil {
		panic(err)
	}
	return &v2.Listener{
		Name: listenerName,
		Address: core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: []listener.FilterChain{{
			Filters: []listener.Filter{{
				Name: util.TCPProxy,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}

//make envoy control plane cache
type SnapshotBuilder struct {
	Version string
	TLS bool
	Listeners []interface{}
	DnsMap map[string]string
}

func (ts SnapshotBuilder) Build() cache.Snapshot {
	listeners := make([]cache.Resource, 0)
	clusters := make([]cache.Resource, 0)
	endpoints := make([]cache.Resource,0)
	routes := make([]cache.Resource,0)
    existsCluster:=make(map[string]string)
	for index,_:= range ts.Listeners {
		config:=ts.Listeners[index]

		l:=config.(*ext.Ingress)
		rules:=l.Spec.Rules
		protocol:=l.Labels["listen.protocol"]
		sport:=l.Labels["listen.port"]
		auth:=l.Labels["listen.auth"]
		fmt.Printf("protocol:%s,port=%s,auth=%s",protocol,sport,auth)
		port,err:=strconv.Atoi(sport)
		if err != nil {
			continue
		}
		protocol=strings.ToUpper(protocol)
		if protocol=="HTTP" {
			fmt.Printf("build http protocol")

			virtualHosts:=make([]route.VirtualHost,0)
			for i,_:= range rules {
				fmt.Printf("###1")
				routeItems :=make([]route.Route,0)
				s:=rules[i]
			    for j,_:= range s.HTTP.Paths {
			    	pathValue:=s.HTTP.Paths[j]
			    	clusterName:=pathValue.Backend.ServiceName
			    	rIP:=ts.DnsMap[clusterName+"."+l.Namespace]
					fmt.Printf("###2")

					if rIP==""{
						fmt.Printf("###x")

						continue
			    	}
					fmt.Printf("###3")

					//不存在集群才构建
			    	if _,ok:=existsCluster[clusterName];!ok{
						fmt.Printf("###4")

						endpoints=append(endpoints,MakeEndpoint(clusterName,rIP,uint32(pathValue.Backend.ServicePort.IntValue())))
			  		   clusters=append(clusters,MakeCluster(Ads,clusterName))
			  	    }
					fmt.Printf("###5")

					routeItems =newRoute(clusterName,auth,pathValue.Path, routeItems)
			    }
				fmt.Printf("###6")

				virtualHosts=append(virtualHosts,MakeVirtualHost(s.Host,s.Host,routeItems))
			}
			if len(virtualHosts)>0 {
				fmt.Printf("###7")

				routes=append(routes,MakeRoute(sport, virtualHosts))
				useSSL:=len(l.Spec.TLS)>0
				listeners = append(listeners, MakeHTTPListener(useSSL,auth,Ads, l.Name, uint32(port), sport))
			}
		}else {
			if l.Spec.Backend != nil {
				s:=l.Spec.Backend
				cluster:=s.ServiceName
				rIP:=ts.DnsMap[s.ServiceName]
				if rIP==""{
					continue
				}
				endpoints=append(endpoints,MakeEndpoint(cluster,rIP,uint32(s.ServicePort.IntValue())))
				clusters=append(clusters,MakeCluster(Ads,cluster))
				listeners=append(listeners,MakeTCPListener(l.Name,uint32(port),cluster))
			}
		}
	}
	out := cache.Snapshot{
		Endpoints: cache.NewResources(ts.Version, endpoints),
		Clusters:  cache.NewResources(ts.Version, clusters),
		Routes:    cache.NewResources(ts.Version, routes),
		Listeners: cache.NewResources(ts.Version, listeners),
	}

	if ts.TLS {
		out.Secrets = cache.NewResources(ts.Version, MakeSecrets(tlsName, rootName))
	}
	error:=out.Consistent()
	if error != nil {
		fmt.Println(error)
	}
	return out
}