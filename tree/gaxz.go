package main

import (
	"sync"
	"github.com/bitly/go-simplejson"
	"strings"
	"encoding/json"
	"fmt"
	"time"
)

type Token struct {
	Type   string `json:"type"`
	Key    string `json:"key"`
	Values []*simplejson.Json `json:"values"`
}

type Parser struct {
	Root   *ParseNode `json:"root"`
	Height int `json:"height"`
	Mu     sync.RWMutex `json:"mu"`
	Update int64 `json:"update"`
}

type ParseNode struct {
	Isroot   bool `json:"isroot"`
	Isleaf   bool  `json:"isleaf"`  // is this a leaf?
	Isparent bool `json:"isparent"` // is this parent, or does this have child(ren)
	Token    Token `json:"token"`
	Parent   *ParseNode `json:"parent"`
	Children []*ParseNode `json:"children"`
}

//针对一个属性多值时
func GetArrayValue(j *simplejson.Json) []string {
	v, ok := j.Interface().(string)
	if ok {
		//字符串
		return []string{v}
	} else {
		//字符串数组
		arr, ok := j.Interface().([]string)
		if ok {
			return arr
		}
		//obj数组
		iarr, ok := j.Interface().([]interface{})
		if ok {
			var sarr []string
			for _, obj := range iarr {
				sarr = append(sarr, obj.(string))
			}
			return sarr
		}
	}
	return []string{}
}
//树中添加一个节点
func (p Parser)AddParseNode(j *simplejson.Json, key, keyword string) {
	ppn := p.GetParseNodeByKeyword(keyword)
	arr := GetArrayValue(j.Get(key))
	for _, v := range arr {
		var pn ParseNode
		//pn.Parent = ppn
		pn.Token.Type = key
		pn.Token.Key = v
		pn.Token.Values = append(pn.Token.Values, j)

		ppn.Children = append(ppn.Children, &pn)
	}
}
//通过查询条件找到节点
func (p Parser)GetParseNodeByKeyword(keyword string) *ParseNode {
	arr := strings.Split(keyword, " ")
	var qtype, qkey []string
	for _, v := range arr {
		v = strings.TrimSpace(v)
		if v != "" {
			tarr := strings.SplitN(v, "=", 2)
			if len(tarr) == 2 {
				qtype = append(qtype, tarr[0])
				qkey = append(qkey, tarr[1])
			}
		}
	}
	_type := strings.Join(qtype, "|")
	_key := strings.Join(qkey, "|")
	return p.GetParseNodeByTypeKey(_type, _key)
}
//通过token的type,key找到节点
func (p Parser)GetParseNodeByTypeKey(qtype, qkey string) *ParseNode {
	if p.Root.Equal(qtype, qtype) {
		return p.Root
	}
	return p.GetParseNode(qtype, qtype, p.Root.Children)
}
//递归找节点
func (p Parser)GetParseNode(qtype, qkey string, pn []*ParseNode) *ParseNode {
	for _, v := range pn {
		if v.Equal(qtype, qkey) {
			return v
		} else {
			tv := p.GetParseNode(qtype, qkey, v.Children)
			if tv != nil {
				return tv
			}
		}
	}
	return nil
}
//////////////////////////////////////////////////
func (pn ParseNode)Equal(qtype, qkey string) bool {
	return pn.Token.Type == qtype && pn.Token.Key == qkey
}

func main() {
	var p Parser
	var pn ParseNode
	pn.Token.Type = "person_phone|person_name"
	pn.Token.Key = "07138888888|廖红伟"
	p.Root = &pn
	for i := 0; i < 3; i++ {
		var pni ParseNode
		pni.Token.Type = "person_qq"
		pni.Token.Key = fmt.Sprint("3243242", i)
		//pni.Parent = pn
		for j := 0; j < 3; j++ {
			var pnj ParseNode
			pnj.Token.Type = "person_email"
			pnj.Token.Key = fmt.Sprint("3243242", i, "@123.com")
			//pnj.Parent = pni
			pni.Children = append(pni.Children, &pnj)
		}
		pn.Children = append(pn.Children, &pni)
	}
	p.Update = time.Now().UnixNano() / 1e6
	byt, _ := json.Marshal(p)
	fmt.Println(string(byt))
}