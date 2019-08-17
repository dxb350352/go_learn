package main

import (
	"os/exec"
	"fmt"
)

func main() {
	err := exec.Command("/bin/bash", "/home/juken/temp/testparam.sh", "host", "dbname", "user", "password", "backuppath").Run()
	if err != nil {
		fmt.Println(err)
	}
}
