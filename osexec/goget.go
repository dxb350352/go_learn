package main
import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println(GoGet("github.com/jackiedong168/chain"))
}

func GoGet(pkg string) error {
	cmd := exec.Command("go", "get", pkg)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
