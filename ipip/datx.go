package main

import (
	"fmt"
	"github.com/ipipdotnet/datx-go"
)

//自己的例子都要报错
func main() {
	//datxpath := "/root/gopath/src/testgo/ipip/17monipdb.datx"
	datxpath := "E:/GOPATH/src/testgo/ipip/17monipdb.datx"
	fmt.Println("city:")
	city, err := datx.NewCity(datxpath)
	if err == nil {
		fmt.Println(city.Find("8.8.8.8"))
		fmt.Println(city.Find("128.8.8.8"))
		fmt.Println(city.Find("255.255.255.255"))
	}
	fmt.Println("dis:")
	dis, err := datx.NewDistrict(datxpath)
	if err == nil {
		fmt.Println(dis.Find("1.12.46.0"))
		fmt.Println(dis.Find("223.255.127.0"))
	}
	fmt.Println("base:")
	bst, err := datx.NewBaseStation(datxpath)
	if err == nil {
		fmt.Println(bst.Find("1.30.6.0"))
		fmt.Println(bst.Find("223.221.121.0"))
		fmt.Println(bst.Find("223.221.121.255"))
	}
}
