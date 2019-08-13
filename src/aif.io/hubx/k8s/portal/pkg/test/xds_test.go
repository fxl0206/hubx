package test

import (
	"testing"
	"net"
	"context"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"google.golang.org/grpc"
	"sync"
	"github.com/prometheus/common/log"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"aif.io/hubx/k8s/portal/pkg/test/resource"

)

var(
	stop = make(chan struct{})
)

func TestXds(t *testing.T)  {
	snapshotCache := cache.NewSnapshotCache(true, Hasher{}, nil)


	signal := make(chan struct{})
	cb := &callbacks{signal: signal}
	server := xds.NewServer(snapshotCache, cb)
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":18000")

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api.RegisterListenerDiscoveryServiceServer(grpcServer, server)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// error handling
		}
	}()
	snapshots := resource.TestSnapshot{
		Xds:              "ads",
		UpstreamPort:     uint32(10001),
		BasePort:         uint32(16002),
		NumClusters:      2,
		Version:"1",
		NumHTTPListeners: 1,
		NumTCPListeners:  1,
		TLS:              false,
	}
	snapshotCache.SetSnapshot("2",snapshots.Generate())
	<-signal

	<-stop
}

type callbacks struct {
	signal   chan struct{}
	fetches  int
	requests int
	mu       sync.Mutex
}

func (cb *callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	//log.WithFields(log.Fields{"fetches": cb.fetches, "requests": cb.requests}).Info("server callbacks")
}
func (cb *callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	log.Debugf("stream %d open for %s", id, typ)
	return nil
}
func (cb *callbacks) OnStreamClosed(id int64) {
	log.Debugf("stream %d closed", id)
}
func (cb *callbacks) OnStreamRequest(int64,  *v2.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.requests++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}
func (cb *callbacks) OnStreamResponse(int64, *v2.DiscoveryRequest, *v2.DiscoveryResponse) {}
func (cb *callbacks) OnFetchRequest(_ context.Context, req *v2.DiscoveryRequest) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.fetches++
	if cb.signal != nil {
		close(cb.signal)
		cb.signal = nil
	}
	return nil
}
func (cb *callbacks) OnFetchResponse(*v2.DiscoveryRequest, *v2.DiscoveryResponse) {}
