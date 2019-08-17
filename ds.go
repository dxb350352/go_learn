package main

import (
	"fmt"
)

type Gooo struct {
	Id int64
}

type TT struct {
	a map[string]string
	b map[string]string
}

func main() {
	f := name()
	fmt.Printf("%p", &f)
	fmt.Println("............")
	deleteTest()
}

func name() Gooo {
	var g Gooo
	fmt.Printf("%p", &g)
	fmt.Println("............")
	return g
}

func deleteTest() {
	var tt TT
	terminalMapCopy := make(map[string]string, 0)
	terminalMapCopy["a"] = "a"
	terminalMapCopy["b"] = "b"
	terminalMapCopy["c"] = "c"
	tt.a = terminalMapCopy
	tt.b = terminalMapCopy
	delete(tt.a, "b")
	for k, v := range tt.a {
		fmt.Println(k, v)
	}
	fmt.Println(".................")
	for k, v := range tt.b {
		fmt.Println(k, v)
	}

}
