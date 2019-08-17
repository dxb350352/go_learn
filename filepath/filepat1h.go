package main
import (
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
)

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	println(path)
	fmt.Println(string(os.PathSeparator))
	fmt.Println(os.Getwd())



	dir := "/fdsa/fdsa/fds/a/../d/fd/"
	fmt.Println("dir:", filepath.Dir(dir))
	dir = filepath.Clean(dir)
	fmt.Println("clean:", dir)
	fmt.Println("dir:", filepath.Dir(dir))
	fmt.Println("base:", filepath.Base(dir))
	fmt.Println("join:", filepath.Join(dir,"ddaa"))

	file_path:="/dd/dd/e.ext"
	fmt.Println(filepath.Abs(file_path))
	pp,ff:=filepath.Split(file_path)
	fmt.Println(pp,ff)
}
