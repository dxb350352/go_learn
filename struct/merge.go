package main

import (
	"reflect"
	"strings"
	"github.com/kr/pretty"
)

type PersonInfo struct {
	PersonId         string `json:"person_id"`
	PersonPhone      string `json:"person_phone"`
	PersonEmail      string `json:"person_email"`
	PersonQQ         string `json:"person_qq"`
	PersonBankcardId string `json:"person_bankcard_id"`
	PersonLoginname  string `json:"person_loginname"`
	PersonRealname   string `json:"person_realname"`
	SchoolName       string `json:"school_name"`
	CompanyName      string `json:"company_name"`
	CompanyFax       string `json:"company_fax"`
	PersonName       string `json:"person_name"`
	PersonNick       string `json:"person_nick"`
	PersonIp         string `json:"person_ip"`
}

func main() {
	var p1 PersonInfo
	p1.PersonId = "dddddd,2"
	arr:=[]PersonInfo{p1}
	var p2 PersonInfo
	p2.PersonId = "33333,2"
	mergeStruct(&(arr[0]), &p2)
	pretty.Println(arr[0])
}

func mergeStruct(p1 interface{}, p2 interface{}) {
	t1 := reflect.ValueOf(p1)
	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}
	t2 := reflect.ValueOf(p2)
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
	}
	for i := 0; i < t1.NumField(); i++ {
		v1 := t1.Field(i).String()
		v2 := t2.Field(i).String()
		if v1 == "" {
			v1 = v2
		} else {
			if v2 != "" {
				v1 += "," + v2
				varr := strings.Split(v1, ",")
				m := map[string]bool{}
				for _, v := range varr {
					m[v] = true
				}
				varr = []string{}
				for k, _ := range m {
					varr = append(varr, k)
				}
				v1 = strings.Join(varr, ",")
			}
		}
		t1.Field(i).SetString(v1)
	}
}
