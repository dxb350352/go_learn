package main
import (
	"io/ioutil"
	"fmt"
	"strings"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(FileReplace("E:/GOPATH_GO/src/github.com/sas/modules/szga1", `"github.com/sas/modules/szga/`, `"github.com/sas/modules/szga1/`))
}

func xiaotest() {
	path1 := "E:/GOPATH_GO/src/testgo/file/replace/applications.test"
	path2 := "E:/GOPATH_GO/src/testgo/file/replace/applications.got"
	cbyte, err := ioutil.ReadFile(path1)
	if err != nil {
		fmt.Println(err.Error())
	}
	content := strings.Replace(string(cbyte), `"github.com/sas/modules/`, `"github.com/sas/modules/fds/`, -1)
	err = ioutil.WriteFile(path2, []byte(content), os.ModeExclusive)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func FileReplace(src, old, new string) error {
	src = filepath.Clean(src)
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(src)
		if err != nil {
			return err
		}
		for _, f := range files {
			err = FileReplace(filepath.Join(src, f.Name()), old, new)
			if err != nil {
				return err
			}
		}
	}else {
		cbyte, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}
		content := strings.Replace(string(cbyte), old, new, -1)
		err = ioutil.WriteFile(src, []byte(content), os.ModeExclusive)
		if err != nil {
			return err
		}
	}
	return nil
}
