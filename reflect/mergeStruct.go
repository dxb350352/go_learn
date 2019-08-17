package main

import (
	"reflect"
	"strings"
	"fmt"
)

type MS struct {
	Name     string
	Children []interface{}
}

func main() {
	ms1 := MS{Name:"name1", Children:[]interface{}{"1", "2"}}
	ms2 := MS{Name:"name2", Children:[]interface{}{"3", "4"}}
	MergeStruct(&ms1, &ms2)
	fmt.Println(ms1)
}
//合并人员信息
func MergeStruct(p1 interface{}, p2 interface{}) {
	t1 := reflect.ValueOf(p1)
	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}
	t2 := reflect.ValueOf(p2)
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
	}
	for i := 0; i < t1.NumField(); i++ {
		switch t1.Field(i).Kind() {
		case reflect.String:
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
}
