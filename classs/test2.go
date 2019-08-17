package main

import (
	"fmt"
	"reflect"
)

type parent struct {
	val int
}

type child struct {
	parent
	num int
}

func main() {

	var c child

	c = child{parent{1}, 2}
	fmt.Println(c.num)
	fmt.Println(c.val)
	var p parent
	fmt.Println(reflect.TypeOf(c).Name())
	fmt.Println(reflect.TypeOf(p))
}
