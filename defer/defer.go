package main


import (
	"fmt"
	"regexp"
	"os"
)

var exp_kuohao = regexp.MustCompile("（.*）")

func main() {
	defer fmt.Println(".............1")
	defer fmt.Println(".............2")
	defer fmt.Println(".............3")
	if len(os.Args) > 0 {
		defer fmt.Println(".............4")
	}
	defer fmt.Println(".............5")
	fmt.Println(tes())
}

func tes() int {
	defer fmt.Println("..............")
	return tes1()
}

func tes1() int {
	fmt.Println("......................test1")
	return 1
}