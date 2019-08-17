package main

import (
	"crypto/tls"
	"github.com/parnurzeal/gorequest"
	"strings"
	"fmt"
)

func main() {
	request := gorequest.New()
	url := "https://192.168.130.240:9007/search/search"
	if strings.HasPrefix(url, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	}
	r := request.Get(url)
	_, body2, errs := r.Query(fmt.Sprintf("startTime=%v", 1486714501000)).
	Query(fmt.Sprintf("endTime=%v", 1491444671000)).Query(fmt.Sprintf("size=%v", 10000)).
	Param("keywords", "__app=ndlp AND __type=events AND c_name=李洪志").End()

	fmt.Println(string(body2),errs)

}
