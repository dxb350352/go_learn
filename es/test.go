package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"bytes"
	"errors"
	"github.com/mattbaird/elastigo/lib"
)
var Es *elastigo.Conn

func init() {
	Es = elastigo.NewConn()
	Es.Hosts = []string{"192.168.130.201"}
	Es.Port = "9200"
}

func main() {
	fmt.Println(Count([]string{`{"hasnew":false}`, `{"phone":"00000000000"}`}))
}

func Count(param []string) (int64, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprint(`{"query":{"bool":{"must":[{"match":{"_index":"attention_attention"}},{"match":{"_type":"attention"}}`))

	for _, val := range param {
		buf.WriteString(fmt.Sprint(`,{"match":`, val, `}`))
	}
	buf.WriteString(`]}}}`)
	body, err := Es.DoCommand("post", "/_count", nil, buf.String())
	if err != nil {
		return 0, err
	}
	j, _ := simplejson.NewJson(body)
	if j.Get("error").Interface() != nil {
		body, _ = j.Get("error").MarshalJSON()
		return 0, errors.New(string(body))
	}
	return j.Get("count").MustInt64(), nil
}
