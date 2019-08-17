package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/mattbaird/elastigo/lib"
	"fmt"
	"encoding/json"
)

var searchFields = []string{
	"person_phone",
	"person_id",
	"person_qq",
	"person_email",
	"person_loginname",
}
var allFields = []string{
	"person_id",
	"person_phone",
	"person_email",
	"person_qq",
	"person_bankcard_id",
	"person_loginname",
	"person_realname",
	"school_name",
	"company_name",
	"company_fax",
	"person_name",
	"person_nick",
	"person_ip",
}

type FieldStatus struct {
	Handled bool `json:"handled"`         //是否已查询
	Score   int  `json:"score"`           //查询到次数
	Set     map[string]bool  `json:"set"` //查询到的__id
	Link    [][]string `json:"link"`      //关系链路
}

func main() {
	ajaj := simplejson.New()
	byt, _ := ajaj.MarshalJSON()
	fmt.Println(string(byt))
	var hits   []elastigo.Hit
	for i := 0; i < 10; i++ {
		j := simplejson.New()
		j.Set("__id", fmt.Sprint("__id_", i))
		for _, v := range searchFields {
			j.Set(v, fmt.Sprint(v, i))
		}
		//fmt.Println(j)
		byt, _ := j.MarshalJSON()
		var rm json.RawMessage = byt
		hits = append(hits, elastigo.Hit{Source:&rm})
	}
	fs := new(FieldStatus)
	arr := []string{"test1"}
	fs.Link = append(fs.Link, arr)
	arr = []string{"test2"}
	fs.Link = append(fs.Link, arr)
	m := getUnique(hits, "person_id", fs)
	for k, v := range m {
		fmt.Println(k, v)
		//for kk, vv := range v {
		//	fmt.Println(k, v, kk, vv)
		//}
	}

}

func getUnique(hits []elastigo.Hit, field string, oldfs *FieldStatus) map[string]map[string]FieldStatus {
	result := make(map[string]map[string]FieldStatus, 0)
	for _, hit := range hits {
		byt, err := hit.Source.MarshalJSON()
		if err != nil {
			return result
		}
		j, err := simplejson.NewJson(byt)
		if err != nil {
			return result
		}
		for _, v := range allFields {
			tj, b := j.CheckGet(v)
			//上次查询的字段拿掉
			if b {
				fmap := result[v]
				if fmap == nil {
					fmap = make(map[string]FieldStatus)
				}
				fs, exist := fmap[tj.MustString()]
				if !exist {
					//当前查询字段不再查询,handled=true
					fs = FieldStatus{Handled:field == v}
					fs.Set = map[string]bool{}
					fs.Link = append(oldfs.Link, []string{})
				}
				//fs.Score = fs.Score + 1
				__id := j.Get("__id").MustString()
				score := len(fs.Set)
				fs.Set[__id] = true
				fs.Score = len(fs.Set)
				//说明有新的数据
				if score < fs.Score {
					fs.Link[len(fs.Link) - 1] = append(fs.Link[len(fs.Link) - 1], __id)
				}

				fmap[tj.MustString()] = fs
				result[v] = fmap
			}
		}
	}
	return result
}