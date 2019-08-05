package main

import (
	"log"
	"os"
	"aif.io/hubx/pkg/root"

	//ensure import portal cmd
	_ "aif.io/hubx/k8s/portal/cmd"
)


func main(){
	if err := root.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}