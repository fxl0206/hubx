package main

import (
	"aif.io/hubx/pkg/root"
	_ "aif.io/hubx/k8s/ipserver/pkg/server"
    _ "aif.io/hubx/k8s/portal/pkg/server"
)


func main(){
	root.Run()
}