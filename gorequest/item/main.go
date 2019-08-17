package main

import (
	"os/exec"
	"fmt"
	"bytes"
)

func main() {
	cmd := exec.Command("/home/juken/gopath/src/test/item/IllegalOnline", "resultUrl", "searchUrl", "123<{value}<321")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out.String())
}
