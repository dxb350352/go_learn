package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jackiedong168/gorequest"
	"strings"
	"regexp"
)

var markreg = regexp.MustCompile(`<span class="mohe-ph-mark".*\n.*\n.*</span>`)
var divreg = regexp.MustCompile(`<div class="gclearfix mh-detail">.*</div>`)

func main() {
	fmt.Println(search("15677246297"))
	fmt.Println(search("4000024005"))
	fmt.Println(search("17092908843"))

}

func search(phone string) (result string) {
	var url = "https://www.so.com/s"
	RequestGo := gorequest.New()
	if strings.HasPrefix(url, "https") {
		RequestGo.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	RequestGo.Get(url)
	RequestGo.Query(fmt.Sprintf("q=%s", phone))
	_, body, _ := RequestGo.End()
	arr := markreg.FindAllString(body, -1)
	result = strings.Join(arr, "")
	if strings.Contains(result, "骚扰") {
		result = "骚扰"
	} else if strings.Contains(result, "诈") {
		result = "诈骗"
	} else {
		result = ""
	}
	return
}
