package main

import (
	"crypto/tls"
	"github.com/bitly/go-simplejson"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
)

func main() {
	Sync_users()
}

func Sync_orgs() {
	request := gorequest.New()
	url := "https://192.168.128.159:9007/v2/sync/orgs"
	if strings.HasPrefix(url, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	}
	json := `{"data":[{"deptid":"2","deptname":"dptndsdsame","pid":""},{"deptid":"221","deptname":"dptndsdsame","pid":"2"}]}`
	data, err := simplejson.NewJson([]byte(json));
	if err != nil {
		fmt.Println(err, ".............")
		return
	}
	request.Post(url)
	_, body, errs := request.SendStruct(data).EndBytes()
	if len(errs) > 0 {
		fmt.Println(errs, ".................errs")
		return
	}
	fmt.Println(string(body))

}
func Sync_users() {
	request := gorequest.New()
	url := "http://192.168.128.159:9005/v2/sync/users"
	if strings.HasPrefix(url, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	}

	json := `{"data":[{"userid":"","user_ic":"","name":"fds2a22a","dept_id":"2","policeno":"6","policetype":"2","mobile":"15987458938"},{"userid":"2","user_ic":"7","name":"fds222a","dept_id":"2","policeno":"","policetype":"2","mobile":"15987458958"}]}`
	data, err := simplejson.NewJson([]byte(json));
	if err != nil {
		fmt.Println(err, ".............")
		return
	}
	request.Post(url)
	_, body, errs := request.SendStruct(data).EndBytes()
	if len(errs) > 0 {
		fmt.Println(errs, ".................errs")
		return
	}
	fmt.Println(string(body))
}
