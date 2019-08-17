package main

import (
	"time"
	"net/url"
	"fmt"
	"reflect"
	"github.com/sas/utils"
)

func main() {
	form := make(url.Values, 0)
	form["id"] = []string{"1"}
	form["name"] = []string{"name"}
	form["created"] = []string{"1483203661123"}
	form["is_parent"] = []string{"1"}
	var obj TestStruct
	Request2Struct(form, &obj)
	fmt.Println(obj)
}

type TestStruct struct {
	Id       int64 `json:"id"`
	Name     string `json:"name"`
	Created  time.Time `json:"created"`
	IsParent bool `json:"is_parent"`
}

func Request2Struct(form url.Values, obj interface{}) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	num := t.NumField()
	for i := 0; i < num; i++ {
		tag := t.Field(i).Tag.Get("json")
		fmt.Println(tag)
		if tag == "" {
			tag = utils.GetDBColumnName(t.Field(i).Name)
		}
		var firstV string
		if len(form[tag]) > 0 {
			firstV = form[tag][0]
		}
		switch t.Field(i).Type.Kind() {
		case reflect.String:
			v.Field(i).SetString(firstV)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.Field(i).SetInt(utils.ParseInt64(firstV))
		case reflect.Bool:
			v.Field(i).SetBool(!(firstV == "false" || firstV == "0" || firstV == ""))
		case reflect.Float32, reflect.Float64:
			v.Field(i).SetFloat(utils.ParseFloat64(firstV))
		case reflect.Struct:
			_, ok := v.Field(i).Interface().(time.Time)
			if ok {
				//只处理时间
				var t time.Time
				var err error
				if utils.IsPositiveInteger(firstV) {
					mics := utils.ParseInt64(firstV)
					t = time.Unix(mics / 1000, mics % 1000)
				} else {
					t, err = time.Parse("2006-01-02 15:04:05", firstV)
					if err != nil {
						t, err = time.Parse("2006-01-02", firstV)
						if err != nil {
							continue
						}
					}
				}
				v.Field(i).Set(reflect.ValueOf(t))
			}

		}
	}
}