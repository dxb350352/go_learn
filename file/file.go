package main
import (
	"io/ioutil"
	"fmt"
	"github.com/sas/utils"
	"os"
)

func main() {
	path := "E:/GOPATH_GO/src/github.com/sas/sas/"
	fmt.Println(utils.CheckFileExist(path))
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("[1]", err)
	}
	version := make(map[string][]string)
	for _, dir := range dirs {
		if dir.IsDir() {
			var files  []string
			vfiles, _ := ioutil.ReadDir(path + dir.Name())
			for _, f := range vfiles {
				if !f.IsDir() {
					files = append(files, f.Name())
				}
			}
			version[dir.Name()] = files
		}
	}
	for k, v := range version {
		fmt.Println(k, v)
	}
	os.Create("e:/ss.s")
	fmt.Println(os.Create("e:/ss.s1/"))
}