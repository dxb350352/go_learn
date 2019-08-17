package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"1":1, "2":2, "3":3}
	mm := mmm(m)
	m["4"] = 4
	for k, v := range mm {
		delete(mm,k)
		fmt.Println(k, v)
	}
	fmt.Println("................")
	for k, v := range mm {
		fmt.Println(k, v)
	}
}

func mmm(m map[string]int) map[string]int {
	m["5"]=5
	return m
}
