package main

import (
	"bytes"
	"testing"
)
import (
	"aif.io/hubx/pkg/root"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
	"os"
	"path/filepath"

	"aif.io/hubx/k8s/portal/api/v1"
	henvoy "aif.io/hubx/k8s/portal/pkg/envoy"
	"aif.io/hubx/k8s/portal/pkg/kube/crd"
	"aif.io/hubx/k8s/portal/pkg/kube/model"
	ext "k8s.io/api/extensions/v1beta1"
	"runtime"
	"time"
)

var(
	stop =make(chan struct{})
)


func TestServer(t *testing.T)  {
	root.SetArgs([]string{"portal","--kubeConfig","test_config"})
	main()
}


func TestDir(t *testing.T){
	_, filename, _, _ := runtime.Caller(1)
	fmt.Println(filename)
	fmt.Println(getExecutePath1())
	fmt.Println(getExecutePath2())
	fmt.Println(getExecutePath3())
	fmt.Println(getExecutePath4())
}

func getExecutePath1() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	return dir
}

func getExecutePath2() string {
	dir, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	exPath := filepath.Dir(dir)
	fmt.Println(exPath)

	return exPath
}

func getExecutePath3() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	return dir
}

func getExecutePath4() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	return dir
}

func  Notify(obj interface{}, event model.Event) error {

	return nil
}
func TestCache(t *testing.T)  {
	cache,err:=crd.MakeKubeConfigController("test_config","",Notify)
	if err != nil {
		fmt.Println(err)
	}
	cache.Run(stop)
	for{
		time.Sleep(1*time.Second)
		configs,err2:=cache.List("listener","default")
		if err2 != nil {
			fmt.Println(err2)
		}
		for i,_:=range configs {
			functionConfig:=configs[i].Spec.(*v1.Listener)
			fmt.Println(functionConfig)
		}
	}

	<- stop

}

func TestBuild(t *testing.T){
	ll:=&ext.Ingress{
	}
	ll.Labels=make(map[string]string)
	ll.Labels["listen.protocol"]="http"
	ll.Labels["listen.port"]="6688"
	ll.Labels["listen.auth"]="2sfwfwf"
	rule:=ext.IngressRule{
		Host:"doc.hubx.site",
		IngressRuleValue:ext.IngressRuleValue{
			HTTP:&ext.HTTPIngressRuleValue{
				Paths:[]ext.HTTPIngressPath{
					{
						Path:    "/",
						Backend: ext.IngressBackend{
							ServiceName:"hubxdoc.docs",
							ServicePort: intstr.FromInt(8899),
						},
					},
				},
			},
		},
	}
	ll.Spec.Rules=[]ext.IngressRule{rule}
	ll.Spec.TLS=[]ext.IngressTLS{
		{
			Hosts:[]string{"doc.hubx.site"},
			SecretName:"doc/test",
		},
	}
	listeners:=[]*ext.Ingress{
		ll,
	}
	llll:=[]interface{}{listeners[0]}
	builder:=henvoy.SnapshotBuilder{DnsMap:map[string]string{"hubxdoc.docs":"10.10.11.1"},Version:"x",Listeners:llll}
	data,err:=json.Marshal(builder.Build())
	if err != nil {
		fmt.Println(err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	fmt.Print(out.String())
}

