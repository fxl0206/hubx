package server

import (
	"aif.io/hubx/k8s/portal/pkg/kube/model"
	"encoding/json"
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"

	"aif.io/hubx/k8s/portal/pkg/modelx"
	"aif.io/hubx/k8s/portal/pkg/tpl"
	"aif.io/hubx/pkg/root"
	"aif.io/hubx/pkg/tools"
	"github.com/spf13/cobra"
	"html/template"
	"k8s.io/api/core/v1"
	"net/http"
	//xdsv1 "aif.io/hubx/k8s/portal/api/v1"
	"sort"
	//"aif.io/hubx/k8s/portal/pkg/kube/crd"
	henvoy "aif.io/hubx/k8s/portal/pkg/envoy"

	"aif.io/hubx/k8s/portal/pkg/envoy/discover"
	"bytes"
	ecache "github.com/envoyproxy/go-control-plane/pkg/cache"

	//"strconv"
)
var(
	kubeConfig string
	grpcPort uint64
	httpPort uint64
	ingressDns string
	stop chan struct{}
	stop2 chan struct{}

	portalCmd = &cobra.Command{
		Use:   "portal",
		Short: "portal service",
		RunE: func(c *cobra.Command, args []string) error {

			//打印参数
			tools.PrintFlags(c.Flags())

			//启动xdsserver
			xdsSnapshotCache := ecache.NewSnapshotCache(true, discover.Hasher{}, nil)
			xdsCallback := &discover.Callbacks{Signal: make(chan struct{}),Cache: xdsSnapshotCache}
			go discover.Start(xdsCallback,grpcPort)

			//初始化k8s原生资源存储
			k8sStore,ingessStore := Start(kubeConfig,"",xdsCallback)

			//启动crd
			//cache,err:=crd.MakeKubeConfigController(kubeConfig,"", xdsCallback.Notify)

			//绑定crd到xdsserver
			//xdsCallback.Store=cache
			/*if err != nil {
				log.Println(err)
				return err
			}
			cache.Run(stop)*/
			xdsCallback.SvcStore =k8sStore
			xdsCallback.IngressStore=ingessStore

			http.HandleFunc("/beat",func(w http.ResponseWriter, r *http.Request){
				fmt.Fprintln(w, "ok")
			})

			http.HandleFunc("/debug",func(w http.ResponseWriter, r *http.Request){
				listeners:=ingessStore.List()

				services:= k8sStore.List()
				dnsMap:=map[string]string{}
				for _,v:=range services{
					svc:=v.(*v1.Service)
					sName:=svc.ObjectMeta.Name+"."+svc.ObjectMeta.Namespace
					if svc.Spec.ClusterIP != ""{
						dnsMap[sName]=svc.Spec.ClusterIP
					}
				}

				builder:=henvoy.SnapshotBuilder{DnsMap:dnsMap,Version:"x",Listeners:listeners}
				data,err:=json.Marshal(builder.Build())
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
				var out bytes.Buffer
				err = json.Indent(&out, data, "", "  ")
				fmt.Fprintln(w,  out.String())
			})

			http.HandleFunc("/debug2",func(w http.ResponseWriter, r *http.Request){
				listeners:=ingessStore.List()
				data,err:=json.Marshal(listeners)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
				var out bytes.Buffer
				err = json.Indent(&out, data, "", "  ")
				fmt.Fprintln(w,  out.String())
			})

			//
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
				services:= k8sStore.List()
				serverInfos:=[]modelx.ServiceInfo{}
				for _,v:=range services{
					svc:=v.(*v1.Service)
					ports:=[]modelx.Port{}
					for _,p:=range svc.Spec.Ports{
						pName:=p.Name
						if pName == ""{
							pName="https"
						}

						uri:=svc.ObjectMeta.Annotations["web.index.uri"]

						if svc.Spec.Type == "NodePort"{
							ports=append(ports,modelx.Port{Name:pName,Target:p.TargetPort.String(),Protocol:string(p.Protocol),Url:fmt.Sprintf("%s://%s:%d%s",pName,ingressDns,p.NodePort,uri)})
						}else if svc.Spec.Type == "LoadBalancer"{
							ports=append(ports,modelx.Port{Name:pName,Target:p.TargetPort.String(),Protocol:string(p.Protocol),Url:fmt.Sprintf("%s://%s:%d%s",pName,ingressDns,p.Port,uri)})
						}else if svc.Spec.ExternalIPs != nil && len(svc.Spec.ExternalIPs)>0{
							var it int
							for it,_= range svc.Spec.ExternalIPs {
								ports=append(ports,modelx.Port{Name:pName,Target:p.TargetPort.String(),Protocol:string(p.Protocol),Url:fmt.Sprintf("%s://%s:%d%s",pName,ingressDns,p.Port,uri)})
							}
							log.Println(it)

						}
					}
					if len(ports)>0 {
						serverInfos=append(serverInfos,modelx.ServiceInfo{ClusterIp:svc.Spec.ClusterIP,Name:svc.Name,ServerIp:ingressDns,Ports:ports})
					}
				}
				/*listeners,err:=cache.List("listener","")
				if err == nil {
					for _,v:= range listeners{
						l:=v.Spec.(*xdsv1.Listener)
						for _,s:=range l.Services {
							ports:=[]modelx.Port{}
							pName:=l.Protocol
							ports=append(ports,modelx.Port{
								Name:pName,
								Target:strconv.FormatUint(uint64(l.Port),10),
								Protocol:l.Protocol,
								Url:fmt.Sprintf("%s://%s:%d%s",pName,ingressDns,
									l.Port,s.Uri)})

							serverInfos=append(serverInfos,modelx.ServiceInfo{
								ClusterIp:"",
								Name:s.Name,
								ServerIp:ingressDns,
								Ports:ports})
						}
					}
				}*/
				t :=template.New("test")

				t,err :=t.Parse(tpl.HTML_TPL)
				if err != nil {
					fmt.Fprintf(w,err.Error())
				}else{
					sort.Slice(serverInfos, func(i, j int) bool {
						res:= serverInfos[i].Name < serverInfos[j].Name
						return res
					})
					err:=t.Execute(w, serverInfos)
					if err != nil {
						fmt.Println(err)
					}
				}
			})
			http.ListenAndServe(tools.ListenPortWrapper(httpPort), nil)
			<-stop
			return nil
		},
	}
)
func init(){
	stop=make(chan struct{})
	stop2=make(chan struct{})

	portalCmd.PersistentFlags().StringVar(&kubeConfig, "kubeConfig", "","k8s config file path")
	portalCmd.PersistentFlags().Uint64Var(&grpcPort, "grpcPort", 8001,"envoy xds server port")
	portalCmd.PersistentFlags().Uint64Var(&httpPort, "httpPort", 8000,"portal http server port")
	portalCmd.PersistentFlags().StringVar(&ingressDns, "ingressDns", "doc.hubx.site","ingress entry dns")

	root.RootCmd.AddCommand(portalCmd)
}
func Start(kubeconfig string,apiServerAddress string,callbacks *discover.Callbacks) (cache.Store,cache.Store){
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags(apiServerAddress, kubeconfig)
	}
	if err != nil {
		panic(err)
	}
	client, err :=kubernetes.NewForConfig(config)
	sharedInformers := informers.NewSharedInformerFactoryWithOptions(client, 3*time.Second, informers.WithNamespace(""))

	svcInformer := sharedInformers.Core().V1().Services().Informer()
	go svcInformer.Run(stop)
	//createCacheHandler(svcInformer,"Service",callbacks)

	ingessInformer := sharedInformers.Extensions().V1beta1().Ingresses().Informer()
	go ingessInformer.Run(stop2)
	createCacheHandler(ingessInformer,"Ingress",callbacks)
	return svcInformer.GetStore(),ingessInformer.GetStore()
}

func createCacheHandler(informer cache.SharedIndexInformer, otype string,callbacks *discover.Callbacks)  {
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println(obj)
				callbacks.Notify(obj,model.EventAdd)
			},
			UpdateFunc: func(old, cur interface{}) {
				fmt.Println(cur)

				callbacks.Notify(cur,model.EventUpdate)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println(obj)

				callbacks.Notify(obj,model.EventDelete)
			},
		})
}
