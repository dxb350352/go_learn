package main

import (
	"fmt"
	"github.com/sas/utils"
	"github.com/jackiedong168/gorequest"
	"github.com/bitly/go-simplejson"
	"errors"
)

func main() {
	req, err := utils.GetSkipValidateRequest("https://192.168.130.240:9005/search/search")
	if err != nil {
		return
	}
	req.Get("https://192.168.130.240:9005/search/search")
	req.Query(fmt.Sprintf("startTime=%v", 0))
	req.Query(fmt.Sprintf("endTime=%v", 1483203665000))
	req.Query(fmt.Sprintf("from=%v", 0))
	//默认一次取100条数据
	req.Query(fmt.Sprintf("size=%v", 100))
	req.Query("keywords=__app=szga AND __type=szga_users_info")
	//获取数据
	fmt.Println(req.QueryData)
	data, err := dDoRequest(req)
	fmt.Println(data, "....................")
	fmt.Println(err, "....................")
}
//执行request并返回结果
func dDoRequest(req *gorequest.SuperAgent) (*simplejson.Json, error) {
	_, body, errs := req.EndBytes()
	if len(errs) > 0 {
		fmt.Println(errs, "....................1")
		return nil, errors.New(fmt.Sprintln(errs))
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println(err, "....................2")
		return nil, errors.New("not return valid json")
	}
	return j2, nil
}