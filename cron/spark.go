package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("/bin/bash", "/home/ws/opt/exec.sh", "/home/ws/work/src/sas/deploy/src/github.com/sas/sas/spark_scripts/terminal_event.py", "1494604800", "0")
	//buf, err := cmd1.Output()
	cmd1.Stdout = os.Stdout
	err := cmd1.Run()
	//fmt.Fprintf(os.Stdout, "Result: %s", buf)
	//err = cmd1.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
