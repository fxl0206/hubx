package server

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/rest"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/informers"
	"time"
	"log"
	"encoding/json"

	"github.com/spf13/cobra"
	"net/http"
	"k8s.io/api/core/v1"
	xdsv1 "aif.io/hubx/k8s/portal/api/v1"
	"sort"
	"html/template"
	"aif.io/hubx/pkg/root"
	"aif.io/hubx/k8s/portal/pkg/tpl"
	"aif.io/hubx/k8s/portal/pkg/modelx"
	"aif.io/hubx/pkg/tools"
	"aif.io/hubx/k8s/portal/pkg/kube/crd"
	henvoy "aif.io/hubx/k8s/portal/pkg/envoy"

	"bytes"
	"aif.io/hubx/k8s/portal/pkg/envoy/discover"
	ecache "github.com/envoyproxy/go-control-plane/pkg/cache"

	"strconv"
)
var(
	kubeConfig string
	grpcPort uint64
	httpPort uint64
	stop chan struct{}
	portalCmd = &cobra.Command{
		Use:   "portal",
		Short: "portal service",
		RunE: func(c *cobra.Command, args []string) error {

			//打印参数
			tools.PrintFlags(c.Flags())

			//初始化k8s原生资源存储
			k8sStore := Start(kubeConfig,"")

			//启动xdsserver
			xdsSnapshotCache := ecache.NewSnapshotCache(true, discover.Hasher{}, nil)
			xdsCallback := &discover.Callbacks{Signal: make(chan struct{}),Cache: xdsSnapshotCache}
			go discover.Start(xdsCallback,grpcPort)

			//启动crd
			cache,err:=crd.MakeKubeConfigController(kubeConfig,"", xdsCallback.Notify)

			//绑定crd到xdsserver
			xdsCallback.Store=cache

			if err != nil {
				log.Println(err)
				return err
			}
			cache.Run(stop)


			http.HandleFunc("/beat",func(w http.ResponseWriter, r *http.Request){
				fmt.Fprintln(w, "ok")
			})

			http.HandleFunc("/debug",func(w http.ResponseWriter, r *http.Request){
				listeners,err:=cache.List("listener","")
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}

				builder:=henvoy.SnapshotBuilder{Version:"x",Listeners:listeners}
				data,err:=json.Marshal(builder.Build())
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
						if svc.Spec.Type == "NodePort"{
							ports=append(ports,modelx.Port{Name:pName,Target:p.TargetPort.String(),Protocol:string(p.Protocol),Url:fmt.Sprintf("%s://%s:%d",pName,"iseex.picp.io",p.NodePort)})
						}else if svc.Spec.Type == "LoadBalancer"{
							ports=append(ports,modelx.Port{Name:pName,Target:p.TargetPort.String(),Protocol:string(p.Protocol),Url:fmt.Sprintf("%s://%s:%d",pName,"iseex.picp.io",p.Port)})
						}
					}
					if len(ports)>0 {
						serverInfos=append(serverInfos,modelx.ServiceInfo{ClusterIp:svc.Spec.ClusterIP,Name:svc.Name,ServerIp:"iseex.picp.io",Ports:ports})
					}
				}
				listeners,err:=cache.List("listener","")
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
								Url:fmt.Sprintf("%s://%s:%d",pName,"iseex.picp.io",
									l.Port)})

							serverInfos=append(serverInfos,modelx.ServiceInfo{
								ClusterIp:"",
								Name:s.Name,
								ServerIp:"iseex.picp.io",
								Ports:ports})
						}
					}
				}
				t :=template.New("test")

				t,err =t.Parse(tpl.HTML_TPL)
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
	portalCmd.PersistentFlags().StringVar(&kubeConfig, "kubeConfig", "","k8s config file path")
	portalCmd.PersistentFlags().Uint64Var(&grpcPort, "grpcPort", 8001,"envoy xds server port")
	portalCmd.PersistentFlags().Uint64Var(&httpPort, "httpPort", 8000,"portal http server port")
	root.RootCmd.AddCommand(portalCmd)
}
func Start(kubeconfig string,apiServerAddress string) cache.Store{
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

	createCacheHandler(svcInformer,"Services")
	return svcInformer.GetStore()
}

func createCacheHandler(informer cache.SharedIndexInformer, otype string)  {
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {

			},
			UpdateFunc: func(old, cur interface{}) {

			},
			DeleteFunc: func(obj interface{}) {

			},
		})
}