package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	j:=simplejson.New()
	j.Set("key",[]byte("abcd"))
	j.Set("a",1)
	j.Set("b","dd")
	fmt.Println(j)
	m:=j.MustMap()
	fmt.Println(m["key"])


	fmt.Println(j.Get("key").Interface().([]byte))
	fmt.Println(j.Get("a").Interface())
	fmt.Println(j.Get("b").Interface())
	j.Set("ddd",1)
	fmt.Println("/.............................")
	fmt.Println(j.Get("ddd").MustFloat64())
	fmt.Println(j.Get("ddd").MustInt())
	fmt.Println(j.Get("ddd").MustInt64())
}
