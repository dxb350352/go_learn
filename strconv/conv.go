package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i64 int64
	var str string
	i64 = 123
	str = strconv.FormatInt(i64, 10)
	fmt.Println(str)
	ii64, err := strconv.ParseInt("_123", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ii64)
	fmt.Println(strconv.FormatInt(int64(0), 10))
}
