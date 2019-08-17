package main
import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	path := "E:/GOPATH_GO/src/github.com/sas/plugins/v3"
	fmt.Println(removeFile(path))
}

func removeFile(path string) error {
	finfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if finfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, v := range files {
			err = removeFile(path + "/" + v.Name())
			if err != nil {
				return err
			}
		}
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}else {
		err = os.Remove(path)
	}
	return err
}