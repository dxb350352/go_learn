package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	j:=simplejson.New()
	j.Set("a","1")
	j.Set("a","2")
	fmt.Println(j)
}
