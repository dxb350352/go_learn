package main

import "fmt"

type Test struct {
	F1 string
	F2 string
}

func main() {
	test1:=Test{F1:"1",F2:"2"}
	test2:=Test{F1:"1",F2:"2"}
	fmt.Println(test1==test2)
	m:=map[Test]int{}
	m[test1]=1
	m[test2]=2
	fmt.Println(m)
}
