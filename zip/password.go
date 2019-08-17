package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("unzip", "-P", "golang", "hello.zip", "-d", "./myconf")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}
