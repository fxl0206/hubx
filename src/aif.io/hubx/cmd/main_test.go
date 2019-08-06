package main

import "testing"
import "aif.io/hubx/pkg/root"


func TestIpServer(t *testing.T)  {
	root.SetArgs([]string{"ipserver","--port","6078"})
	main()
}

func TestPortalServer(t *testing.T)  {
	root.SetArgs([]string{"portal","--kubeConfig","test_config"})
	main()
}