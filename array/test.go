package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	arr:=[]string{"a"}
	j:=simplejson.New()
	j.Set("name",arr)
	fmt.Println(j.Get("name").Interface().([]string))
}
