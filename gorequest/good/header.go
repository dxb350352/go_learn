package main

import (
	"github.com/jackiedong168/gorequest"
	"fmt"
)

func main() {
	recipient := "http://192.168.130.201:9000/exact/getnameid?phone=13242503108"
	method := "get"
	RequestGo := gorequest.New()
	//RequestGo.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	if method == "get" {
		//get
		RequestGo.Get(recipient)
		RequestGo.Header["X-Ssl-Client-S-Dn"] = "ST=\xE5\x85\xA8\xE5\x9B\xBD/L=\xE5\x85\xA8\xE7\x9C\x81/CN=\x30\x30\x30\x30\x30\x30\x30\x30\x30\x30\x30\x66\x6C\x6A\x73\x61\x6C\x6A\x6A\x6B\x6C\x64\x73\x6C\x6A\x6B"

		fmt.Println(recipient, "................get")
	} else {
		//post
		RequestGo.Post(recipient)
		RequestGo.Header["X-Ssl-Client-S-Dn"] = "ST=\xE5\x85\xA8\xE5\x9B\xBD/L=\xE5\x85\xA8\xE7\x9C\x81/CN=\x30\x30\x30\x30\x30\x30\x30\x30\x30\x30\x30\x66\x6C\x6A\x73\x61\x6C\x6A\x6A\x6B\x6C\x64\x73\x6C\x6A\x6B"
		fmt.Println(recipient, "................post")
	}
	resp, body, errs := RequestGo.End()
	fmt.Println(resp, body, errs)

	//req:=gorequest.New()
	//req.Get("http://192.168.128.198:9001/public/img/favicon.png")
	//req.Set("Referer","192.168.130.181")
	//fmt.Println(req.Header["Referer"])
	//fmt.Println(req.AsCurlCommand())
	//req.EndBytes()

	//gorequest.New().
	//	Get("http://192.168.128.198:9001/public/img/favicon.png").
	//	Set("Referer", "192.168.130.181").
	//	End()
}
