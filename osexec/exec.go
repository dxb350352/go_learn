package main
import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	cmd := exec.Command("revel", "run github.com/sas/sasslvpn_web prod 9000")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
//需要root权限时
func Root() {
	cmd := exec.Command("/bin/bash", "/home/juken/restart.sh")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	cmd.Stdout.Write([]byte("root_password"))
}