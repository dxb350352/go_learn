package main

import (
	"strings"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	j:=simplejson.New()
	j.Set("person2_qq","qq")
	var arr []string
	if tj, exist := j.CheckGet("person_qq"); exist {
		arr = strings.Split(tj.MustString(), ",")
	}
	fmt.Println(len(arr))
	arr = append(arr, "a")
	fmt.Println(strings.Join(arr, ","))
}
