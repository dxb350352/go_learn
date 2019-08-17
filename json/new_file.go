package main

import (
	"encoding/json"
	"fmt"
	js "github.com/bitly/go-simplejson"
	"time"
	"github.com/henrylee2cn/pholcus/common/simplejson"
)

func main() {
	fmt.Println(simplejson.NewJson([]byte(fmt.Sprint([]string{"123","323ds"}))))
	str := "[{\"panelId\":1}]"
	fmt.Println(str)
	m1, err := js.NewJson([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m1)
	s, err := json.Marshal(m1)
	fmt.Println(string(s))
	tim := time.Now()
	fmt.Println(tim)
	s,err =json.Marshal(tim)
	fmt.Println(string(s))
	
}
