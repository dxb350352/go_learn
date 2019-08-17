package main

import (
	"os"
	"log"
	"fmt"
	"strings"
	"encoding/xml"
	"github.com/bitly/go-simplejson"
	"encoding/base64"
)

var KeyChange = map[string]string{

}

func main() {
	file, err := os.Open("E:/GOPATH/src/testgo/xml/1560010_ACCOUNT_3_44030015152836225403540_0541_V2.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	var j *simplejson.Json
	var ii int64
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		if err != nil {
			log.Fatal(err)
		}
		ii++
		switch token := t.(type) {
		case xml.StartElement:
			// 处理元素开始（标签）
			name := token.Name.Local
			if name == "DATA" {
				j = simplejson.New()
				continue
			}
			if name == "ITEM" {
				m := map[string]string{}
				for _, attr := range token.Attr {
					m[attr.Name.Local] = attr.Value
				}
				fmtstr := strings.ToLower(m["fmt"])
				key := strings.ToLower(m["key"])
				val := m["val"]
				if fmtstr == "" {
					j.Set(key, strings.ToLower(val))
				} else if fmtstr == "base64" {
					byt, err := base64.StdEncoding.DecodeString(val)
					if err != nil {
						continue
					}
					val = string(byt)
					j.Set(key, strings.ToLower(val))
				}
			}

		case xml.EndElement:
			// 处理元素结束（标签）
			name := token.Name.Local
			if name == "DATA" && j != nil {
				byt, _ := j.MarshalJSON()
				fmt.Println(string(byt))
			}
		default:
		}
	}
	fmt.Println(ii)
}
