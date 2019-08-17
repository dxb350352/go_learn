package main

import (
	"fmt"
	"strings"
)

func main() {
	//test()
	test1()
}
func test1()  {
	str1 := "aabbccbbaa"
	fmt.Println(strings.TrimPrefix(str1,"aaa"))
	fmt.Println(strings.TrimLeft(str1,"aaa"))
	result:=strings.TrimFunc(str1, func(ru rune) bool{
		fmt.Println(string(ru))
		return ru==rune('a')
	})
	fmt.Println(result)
}
func test()  {
	str1 := "aabbccbbaa"
	fmt.Println(strings.Trim(str1, "b"))
	fmt.Println(strings.Trim(str1, "a"))
	fmt.Println(strings.Trim(str1, "ab"))
	fmt.Println(strings.Trim(str1, "ba"))
}
