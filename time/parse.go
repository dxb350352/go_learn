package main

import (
	"time"
	"regexp"
	"strings"
	"fmt"
)

var NumberReg = regexp.MustCompile(`\d+`)

func main() {
	expire := "2099-09-09T00:00:00Z"
	arr := NumberReg.FindAllString(expire, 3)
	expire = strings.Join(arr, "-")
	cst, err := time.LoadLocation("Asia/Shanghai")
	fmt.Println(cst, err)
	t, err := time.ParseInLocation("2006-01-02", expire, cst)
	fmt.Println(t, err)
	fmt.Println(t.Before(time.Now()))
}
