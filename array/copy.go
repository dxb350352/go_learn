package main

import "fmt"

func main() {
	//guessingGame()
	//ssss()
	var arr = []string{"a", "b"}
	var a []string = make([]string, len(arr))
	fmt.Println(copy(a, arr))
	fmt.Println(a)
	fmt.Println(arr)
	var str = "ab"
	var s []byte = make([]byte, len(str))
	fmt.Println(copy(s, str))
	fmt.Println(string(s))

	var aa = []int{1, 2, 3}
	bb := aa
	bb = append(bb[:1], bb[2:]...)
	fmt.Println(aa, bb)

	aa = []int{1, 2, 3}
	bb = make([]int, len(aa))
	copy(bb, aa)
	fmt.Println(aa, bb)
}
