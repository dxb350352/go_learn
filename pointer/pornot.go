package main
import "fmt"


type ss struct {
	id int
}

func (s *ss)add() {
	s.id++
}
func (s ss) add1() {
	s.id++
}
func main() {
	test2()
	fmt.Println("--------------------------")
	test1()
}

func test2() {
	s := ss{id:1}
	s.add()
	fmt.Println(s)
	s.add1()
	fmt.Println(s)
}

func test1() {
	var a Integer = 1
	p := &a
	fmt.Println(1, a, &p)
	a.Add()
	p = &a
	fmt.Println(2, a, &p)
	a.Add1()
	p = &a
	fmt.Println(3, a, &p)
}

type Integer int

func (a *Integer) Add() {
	fmt.Println(a)
	*a++
}
func (a Integer) Add1() {
	p := &a
	fmt.Println(a, &p)
	a++
}