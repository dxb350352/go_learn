package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, e := ioutil.ReadDir("E:/GOPATH/src/testgo/file/gaxz")
	os.Create("E:/GOPATH/src/testgo/file/gaxz/test.txt")
	fmt.Print(f, e)
	for _, v := range f {
		fmt.Println(v.Name())
	}
	//os.MkdirAll("d:/1/2/3/", os.ModeDir)
}
