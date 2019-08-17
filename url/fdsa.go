package main

import (
	"net/url"
	"fmt"
)

func main() {
	p:="/1/21?ss=22"
	u, err := url.Parse(p)
	if err != nil {
		fmt.Println(err.Error(),"eiroror")
	}
	fmt.Println(u,"....")
	fmt.Println(u.RequestURI())
	fmt.Println(u.Query().Get("ss"))
	fmt.Println(url.PathEscape("走你"))
	fmt.Println(url.PathUnescape("%E8%B5%B0%E4%BD%A0"))
}
