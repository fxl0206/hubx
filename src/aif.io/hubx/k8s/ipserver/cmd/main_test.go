package main

import "testing"
import "aif.io/hubx/pkg/root"


func TestServer(t *testing.T)  {
	root.SetArgs([]string{"ipserver","--port","6078"})
	main()
}