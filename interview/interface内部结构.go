package main

import (
	"fmt"
)
//https://my.oschina.net/u/553243/blog/1478739
type Peoples interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() Peoples {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
