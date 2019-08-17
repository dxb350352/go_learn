package main

import (
	"github.com/smallnest/goreq"
	"github.com/bitly/go-simplejson"
	"crypto/tls"
	"fmt"
	"errors"
	"time"
)

var (
	repoType string = "fs"//仓库类型
	RepoName string = "esbackup"//仓库名称
	repoPath = "/opt/esbackup"//仓库路径
	Indices = "sas_*"
)

func main() {
	testCreate()
}

func testCreate() {
	request := goreq.New()
	request.Get(getURL() + "/esbackup/createrepo")
	request.Param("repoType", repoType)
	request.Param("name", RepoName)
	request.Param("path", repoPath)
	js, err := DoRequest(request)
	fmt.Println(js, err)
}

func testBackup() {
	request := goreq.New()
	request.Get(getURL() + "/esbackup/createshapshot")
	request.Param("repo", RepoName)
	request.Param("name", time.Now().Format("20060102-150405"))
	request.Param("indices", Indices)
	js, err := DoRequest(request)
	fmt.Println(js, err)
}

func testList() {
	request := goreq.New()
	request.Get(getURL() + "/esbackup/getshapshotstatus")
	request.Param("repo", RepoName)
	js, err := DoRequest(request)
	fmt.Println(js, err)
}

func getURL() string {
	return "192.168.132.175:9012"
}

//执行request并返回结果
func DoRequest(request *goreq.GoReq) (*simplejson.Json, error) {
	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, body, errs := request.End()
	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintln(errs))
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return nil, errors.New("not return valid json")
	}
	if j2.Get("status").MustInt() != 200 {
		return nil, errors.New(body)
	}
	return j2.Get("data"), nil
}