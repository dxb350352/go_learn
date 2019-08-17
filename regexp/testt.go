package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := `11:22:33:44`
	exp := `\d+(?=((?!\d+).)*$)`
	reg := regexp.MustCompile(exp)
	result := reg.FindAllStringSubmatch(str, -1)
	for _, v := range result {
		fmt.Println(v)
	}
}
