package main

import "fmt"

type AAA struct {
	name string
}

func main() {
	m := map[string]*AAA{
		"a":&AAA{name:"a"},
		"b":&AAA{name:"b"},
	}

	v := m["a"]
	delete(m, "a")
	fmt.Println(v)
}
