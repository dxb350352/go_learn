package main

import "fmt"

type TestStruct struct {
	Age int
}

func main() {
	////map是地址传递--------------就Map不一样
	//testmap()
	////struct是值传递
	testStruct()
	////array是值传递
	//testArray()
}
func testArray() {
	var arr = []int{1, 2, 3}
	tarr := testArrayF(arr)
	fmt.Println(arr)
	fmt.Println(tarr)
}
func testArrayF(arr []int) []int {
	arr = append(arr, 4)
	return arr
}
func testStruct() {
	var ts TestStruct
	tss := testStructF(ts)
	fmt.Println(ts)
	fmt.Println(tss)
	ttss:=tss
	ttss.Age=3
	fmt.Println(tss)
	fmt.Println(ttss)
}

func testStructF(ts TestStruct) TestStruct {
	ts.Age = 2
	return ts
}

func testmap() {
	m := map[string]int{"1":1, "2":2, "3":3}
	mm := testmapF(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, v := range mm {
		fmt.Println(k, v)
	}
}

func testmapF(m map[string]int) map[string]int {
	m["5"] = 5
	return m
}