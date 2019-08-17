package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := `root         1  0.0  0.0  19364  1540 ?        Ss   Jul12   0:00 /sbin/init  dddd sss`
	exp := `(\S)+(\s)*`
	reg := regexp.MustCompile(exp)
	result := reg.FindAllStringSubmatch(str, -1)
	for _, v := range result {
		fmt.Println(v)
	}
}
