package main

import (
	"fmt"
	"crypto/tls"
	"github.com/jackiedong168/gorequest"
	"time"
	"github.com/bitly/go-simplejson"
	"errors"
)

var SearchUrl string = "https://192.168.130.247:9005/search/search"
var pageSize int64 = 100

func main() {

	var i, total int64 = 1, pageSize
	for ; i * pageSize <= total; i++ {
		fmt.Println("i", i)
		req := GetRequest(i)
		data, err := DoRequest(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		total = data.Get("total").MustInt64()
		if total == 0 {
			fmt.Println("00000000000")
			return
		}
		fmt.Println("total:", total)
		data = data.Get("hits")
		for j := 0; ; j++ {
			jd := data.GetIndex(j)
			if jd.Interface() == nil {
				break
			}

		}
	}
}

func DoRequest(req *gorequest.SuperAgent) (*simplejson.Json, error) {
	_, body, errs := req.EndBytes()
	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintln(errs))
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return nil, errors.New("not return valid json")
	}
	return j2, nil
}

func GetRequest(page int64) *gorequest.SuperAgent {
	req := gorequest.New()
	req.Transport.DisableKeepAlives = false
	req.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req.Get(SearchUrl)
	req.Query(fmt.Sprintf("startTime=%v", 0))
	req.Query(fmt.Sprintf("endTime=%v", time.Now().Unix() * 1000))
	req.Query(fmt.Sprintf("from=%v", (page - 1) * pageSize))
	req.Query(fmt.Sprintf("size=%v", pageSize))
	req.Query("keywords=__app=sas_v3 AND __type=t_log_test")
	return req
}

func write2File(j *simplejson.Json)  {

}