package main

import (
	"fmt"
)

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

func test1() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	arr = append(arr[:3], arr[4:]...)
	fmt.Println(arr)
}
func test2() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	for i, v := range arr {
		if v == 3 || v == 4 {
			arr = append(arr[:i], arr[i + 1:]...)
			continue
		}
	}
	fmt.Println(arr)
}
func test3() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	for i, v := range arr {
		fmt.Println(i, v)
		//可以删，但不能再循环了
		arr = append(arr[:i], arr[i + 1:]...)
	}
}

func test4() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	for i := 0; i < len(arr);i++ {
		if arr[i] == 5 || arr[i] == 6 {
			arr = append(arr[:i], arr[i + 1:]...)
			i--
		}

	}
	fmt.Println(arr)
}