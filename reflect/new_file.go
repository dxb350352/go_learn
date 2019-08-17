package main

import (
	"reflect"
	"fmt"
	"strings"
	"time"
)
type ss struct {
	Id   int    `json:"id"colName:"ID"`
	Name string `json:"name"colName:"名称"`
	TT   time.Time
}
func main() {
//	a := models.BMTrigger{}
	v := ss{Id:123, Name:"fdsa", TT:time.Now()}
	m := make(map[string]interface{})
	fmt.Println(v)
	GetObjectKV(&v, m)
	for kk, vv := range m {
		fmt.Println(kk, vv)
	}
	mm := GetObjectKT(ss{Id:123, Name:"gogo"}, "json")
	for kk, vv := range mm {
		fmt.Println(kk, vv, "...")
	}
}

func GetObjectKV(o interface{}, m map[string]interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	//所有字段加上类名做为前缀
	preName := t.Name() + "."
	var temp interface{}
	for i := 0; i < t.NumField(); i++ {
		temp = v.FieldByName(t.Field(i).Name).Interface()
		if reflect.TypeOf(temp).Kind() == reflect.Struct {
			myt, ok := temp.(time.Time)
			if ok {
				m[strings.ToUpper(preName + t.Field(i).Name)] = myt.Format("2006-01-02 15:04:05")
				continue
			}
			tt := reflect.TypeOf(temp)
			vv := reflect.ValueOf(temp)
			if tt.Kind() == reflect.Ptr {
				tt = tt.Elem()
				vv = vv.Elem()
			}
			for ii := 0; ii < tt.NumField(); ii++ {
				m[strings.ToUpper(preName + tt.Field(ii).Name)] = vv.FieldByName(tt.Field(ii).Name).Interface()
			}
		}else {
			m[strings.ToUpper(preName + t.Field(i).Name)] = temp
		}
	}
}

func GetObjectKT(o interface{}, tag string) map[string]string {
	m := make(map[string]string)
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	preName := t.Name() + "."
	for i := 0; i < t.NumField(); i++ {
		m[strings.ToUpper(preName + t.Field(i).Name)] = t.Field(i).Tag.Get(tag)
	}
	return m
}