package main

import (
	"fmt"
	"github.com/ipipdotnet/datx-go"
	"strings"
)

func main() {
	datxpath := "E:/GOPATH/src/testgo/ipip/17monipdb.datx"
	city, err := datx.NewCity(datxpath)
	if err == nil {
		find_print_ip(city, "182.71.161.154")
		find_print_ip(city, "103.62.236.198")
		find_print_ip(city, "209.126.74.124")
		find_print_ip(city, "115.196.95.6")
		find_print_ip(city, "218.14.22.77")
	}
}

func find_print_ip(city *datx.City, ip string) {
	arr, _ := city.Find(ip)
	fmt.Println(strings.Join(arr, ","), len(arr))
}
