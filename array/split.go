package main

import (
	"strings"
	"fmt"
)

func main() {
	split2arr("")
	split2arr("/")
	split2arr("/fdsa")
	split2arr("/fdas/ffg")
	join("123_3_fdkljklfdsa")
	join("123_3_fdkljklfdsa_fdsa")
}

func split2arr(str string) {
	arr := strings.Split(str, "/")
	fmt.Println(len(arr), strings.Count(str, "/"), "..................")
	for i, v := range arr {
		fmt.Println(str, i, v)
	}
}

func join(str  string) {
	arr := strings.Split(str, "_")
	fmt.Println(strings.Join(arr[2:], "_"))
}