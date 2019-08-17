package main

import (
	"sort"
	"fmt"
)

func main() {
	a:=[]string{"a","c","b"}
	sort.Strings(a)
	fmt.Println(a)
}
