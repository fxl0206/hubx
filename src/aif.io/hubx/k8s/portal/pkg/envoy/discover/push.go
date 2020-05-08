package discover

import (
	"k8s.io/api/core/v1"
	"sync"
	"context"
	"aif.io/hubx/k8s/portal/pkg/kube/model"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/prometheus/common/log"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	henvoy "aif.io/hubx/k8s/portal/pkg/envoy"
	kcache "k8s.io/client-go/tools/cache"

	"fmt"
	"strconv"
	"time"
)

type Callbacks struct {
	Signal   chan struct{}
	fetches  int
	requests int
	version int
	mu       sync.Mutex
	timer *time.Timer
	Store  model.ConfigStoreCache
	K8sStore kcache.Store

	Cache cache.SnapshotCache
}

func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	//log.WithFields(log.Fields{"fetches": cb.fetches, "requests": cb.requests}).Info("server callbacks")
}
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	log.Debugf("stream %d open for %s", id, typ)
	return nil
}
func (cb *Callbacks) OnStreamClosed(id int64) {
	log.Debugf("stream %d closed", id)
}
func (cb *Callbacks) OnStreamRequest(int64,  *v2.DiscoveryRequest) error {
	fmt.Println("OnStreamRequest")
	cb.Signal <- struct{}{}
	return nil
}
func (cb *Callbacks) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {
}
func (cb *Callbacks) OnFetchRequest(_ context.Context, req *v2.DiscoveryRequest) error {
	return nil
}
func (cb *Callbacks) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {
}

func (cb *Callbacks) Push() error{
	if cb.Store == nil || cb.Cache == nil {
		fmt.Println("push do nothing")
		return nil
	}

	for _,key:= range cb.Cache.GetStatusKeys(){
		listeners,err:=cb.Store.List("listener","")
		if err != nil {
			return err
		}
		services:= cb.K8sStore.List()
		dnsMap:=map[string]string{}
		for _,v:=range services{
			svc:=v.(*v1.Service)
			sName:=svc.ObjectMeta.Name+"."+svc.ObjectMeta.Namespace
			if svc.Spec.ClusterIP != ""{
				dnsMap[sName]=svc.Spec.ClusterIP
			}
		}
		builder:=henvoy.SnapshotBuilder{DnsMap:dnsMap,Version:strconv.Itoa(cb.version),Listeners:listeners}
		fmt.Println(fmt.Sprintf("push id=%s cache",key))
		ss:=builder.Build()
		cb.Cache.SetSnapshot(key,ss)
	}
	return nil
}

func (cb *Callbacks) Notify(obj interface{}, event model.Event) error {
	cb.version++
	fmt.Println(obj,event)
	cb.Signal <- struct{}{}
	return nil
}

func (cb *Callbacks) Loop(){
	for{
		cb.loopPusher()
	}
}

func (cb *Callbacks) loopPusher(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("#push fail  reson: ",err)
		}
	}()
	select {
	case <-cb.Signal:
		if cb.timer!= nil {
			cb.timer.Reset(1*time.Second)
		}else{
			cb.timer=time.AfterFunc(1*time.Second, func() {
				cb.Push()
			})
		}
	}
}
