package main

import (
	"github.com/mattbaird/elastigo/lib"
	"fmt"
)

var Es *elastigo.Conn

func init() {
	Es = elastigo.NewConn()
	Es.Hosts = []string{"192.168.130.201"}
	Es.Port = "9203"
}
func main() {
	urlstr := "/cluster_one:*,cluster_two:*,*/xx2,xx1,xx3/_search?q=*&pretty&ignore_unavailable=true"
	body, err := Es.DoCommand("POST", urlstr, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
