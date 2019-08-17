package main

import (
	"time"
	"crypto/tls"
	"github.com/parnurzeal/gorequest"
	"github.com/sas/utils"
	"fmt"
)

func main() {
	test1()
}

func test1()  {
	rr := gorequest.New().Timeout(time.Minute).Get("https://192.168.128.250:50089/rest/addStrategy.json")
	rr.Param("Action", "404")
	rr.Param("ClientIP", "192.168.128.51")
	rr.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	json, err := utils.DoRequest(rr)
	fmt.Println(json, err)
}

func test2()  {
	rr := gorequest.New().Timeout(time.Minute).Post("https://192.168.130.240:9007/Terminals/V1_statuses")
	rr.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	json, err := utils.DoRequest(rr)
	fmt.Println(json, err)
}