package main

import (
	"aif.io/hubx/pkg/root"
	_ "aif.io/hubx/k8s/ipserver/pkg/server"
)

func main(){
	root.Run()
}
