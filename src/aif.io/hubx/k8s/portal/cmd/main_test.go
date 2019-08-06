package main

import "testing"
import (
	"aif.io/hubx/pkg/root"
	"fmt"
	"path/filepath"
	"os"
	"runtime"
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