package main

import (
	"fmt"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/kr/pretty"
)

type Test struct {
	AAA string `json:"a"`
}

func main() {
	var test Test
	m := map[string]string{"a":"1"}
	byt,_:=json.Marshal(m)
	json.Unmarshal(byt,&test)
	fmt.Println(test)
	fordel()
}

func fordel()  {
	var arr []*simplejson.Json
	for i:=0;i<10;i++{
		j:=simplejson.New()
		j.Set("__id",fmt.Sprint(i))
		arr=append(arr,j)
	}
	j:=simplejson.New()
	j.Set("data",arr)
	arrr:=j.Get("data").Interface().([]*simplejson.Json)
	for _,v:=range arrr{
		id := v.Get("__id").MustString()
		if id == "" {
			pretty.Println("..............ddddddddddddddddddd")
		}
		v.Del("__id")
	}
}