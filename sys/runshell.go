package main
import (
	"os/exec"
	"fmt"
)

func main() {
	Ipconfig()
}

func Bash() {
	cmd := exec.Command("/bin/bash", "xx.sh")
	err := cmd.Run()
	print(err)
}

func Ipconfig() {
	cmd := exec.Command("ipconfig", "-all")
	b, err := cmd.Output()
	print(err)
	fmt.Println(string(b))
}

func print(i interface{}) {
	if i != nil {
		fmt.Println(i)
	}
}