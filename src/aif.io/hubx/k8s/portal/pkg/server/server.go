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
	"github.com/spf13/cobra"
	"net/http"
	"k8s.io/api/core/v1"
	"sort"
	"html/template"
	"aif.io/hubx/pkg/root"
	"aif.io/hubx/k8s/portal/pkg/tpl"
	"aif.io/hubx/k8s/portal/pkg/modelx"
	"aif.io/hubx/pkg/tools"
)
var(
	kubeConfig string
	stop chan struct{}
	portalCmd = &cobra.Command{
		Use:   "portal",
		Short: "portal service",
		RunE: func(c *cobra.Command, args []string) error {
			tools.PrintFlags(c.Flags())
			store:= Start(kubeConfig,"")

			http.HandleFunc("/beat",func(w http.ResponseWriter, r *http.Request){
				fmt.Fprintln(w, "ok")
			})
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
				services:=store.List()
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
			http.ListenAndServe(":8000", nil)
			<-stop
			return nil
		},
	}
)
func init(){
	stop=make(chan struct{})
	portalCmd.PersistentFlags().StringVar(&kubeConfig, "kubeConfig", "","k8s config file path")
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