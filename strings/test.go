package main

import (
	"fmt"
	"strings"
)

func main() {
	a:=[]string{"a","b","c","d"}
	fmt.Println(strings.Join(a[:2]," "))
}
