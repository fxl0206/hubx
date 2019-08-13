package crd

import (
	"aif.io/hubx/k8s/portal/pkg/kube/model"
	"github.com/hashicorp/go-multierror"
	"time"
	"aif.io/hubx/k8s/portal/pkg/kube"
	_ "aif.io/api/networking/v1"
	_ "aif.io/hubx/k8s/portal/api/v1"
)

func MakeKubeConfigController(kubeCfgFile string ,domainSuffix string,xupdater kube.Handler) (model.ConfigStoreCache, error) {
	configClient, err := NewClient(kubeCfgFile, "", model.IstioConfigTypes, domainSuffix)
	if err != nil {
		return nil, multierror.Prefix(err, "failed to open a config client.")
	}

	if err = configClient.RegisterResources(); err != nil {
		return nil, multierror.Prefix(err, "failed to register custom resources.")
	}

	return NewController(configClient, kube.ControllerOptions{
		WatchedNamespace:"",
		ResyncPeriod:2*time.Second,
		DomainSuffix:domainSuffix,
		Xupdater:xupdater,
		Stop:make(chan struct{}),
	}), nil
}