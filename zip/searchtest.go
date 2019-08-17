package main

import (
	"github.com/sas/gkzx/app/models"
	"time"
	"github.com/sas/utils"
	"fmt"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"net/url"
	"strconv"
)

func main() {
	list()
}

func create() {
	var log models.SzgaSysLog
	log.CreateTime = time.Now()
	log.Ip = "127.0.0.1"
	log.Login = "管理"
	log.Message = "验证失败"
	log.Path = "szga_user/create"
	log.Result = "失败"
	log.CreateTimeUnix = log.CreateTime.Unix() * 1000
	js, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err, "..............1")
	}
	sjs, err := simplejson.NewJson(js)
	if err != nil {
		fmt.Println(err, "..............2")
	}
	err = utils.IndexingBranch("https://192.168.130.240:9007/search/bulk", "szga", "szga_sys_log", []*simplejson.Json{sjs})

	fmt.Println(err, ".............3")
}
func list2() {//这个不行，我了个去
	login := ""
	result := ""
	url_ := "https://192.168.130.240:9007/search/search"
	path, err := url.Parse(url_)
	if err != nil {
		fmt.Println(err)
	}
	path.Query().Set("from", strconv.Itoa(0))
	path.Query().Set("size", strconv.Itoa(10))
	path.Query().Set("startTime", strconv.FormatInt(time.Now().Add(-time.Hour * 200).Unix() * 1000, 10))
	path.Query().Set("endTime", strconv.FormatInt(time.Now().Unix() * 1000, 10))
	path.Query().Set("app", "szga")
	path.Query().Set("_type", "sys_log")
	params := "__type=szga_sys_log"
	if login != "" {
		params += " AND login=" + login
	}
	if result != "" {
		params += " AND result=" + result
	}
	path.Query().Set("keywords",params)
	fmt.Println(path.Query()["app"],".......")
	fmt.Println(path.RequestURI())
	fmt.Println(utils.GetDataFromSeachUrl(url_, path.Query().Encode()))
}
func list() {
	login := "管理"
	result := ""
	url_ := "https://192.168.130.240:9007/search/search"
	params := fmt.Sprint("from=", 0, "&size=", 10, "&startTime=", time.Now().Add(-time.Hour * 200).Unix() * 1000, "&endTime=", time.Now().Unix() * 1000)
	params += fmt.Sprint("&app=szga&_type=sys_log")
	params += "&keywords=__type=szga_sys_log"
	if login != "" {
		params += " AND login=" + login
	}
	if result != "" {
		params += " AND result=" + result
	}
	fmt.Println(url_)
	fmt.Println(params)
	v, err := url.ParseRequestURI(url_ + "?" + params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Query().Encode())
	fmt.Println(utils.GetDataFromSeachUrl(url_, params))
	fmt.Println(utils.GetDataFromSeachUrl(url_, v.Query().Encode()))
}
