package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
	"encoding/json"
)

func main() {
	var data interface{}
	data, _ = simplejson.NewJson([]byte(`{"a":1}`))
	var request string
	switch v := data.(type) {
	case string:
		request = v
		fmt.Println("string")
	case []byte:
		request = string(v)
		fmt.Println("[]byte")
	default:
		byt, err := json.Marshal(v)
		if err == nil {
			request = string(byt)
		}
	}
	fmt.Println(request)
}
