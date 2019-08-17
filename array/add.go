package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	params := []string{"a"}
	params = append(params, strings.Split("b c   d", " ")...)
	multiParams(params...)
}

func multiParams(args ...string) {
	for i, v := range args {
		fmt.Println(i, "----", v)
	}

}