package main

import (
	"fmt"
	"github.com/wangtuanjie/ip17mon"
)

func init() {
	if err := ip17mon.Init("E:/GOPATH/src/testgo/ipip/17monipdb.dat"); err != nil {
		panic(err)
	}
}

func main() {
	loc, err := ip17mon.Find("116.228.111.18")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(loc)
}
