package main
import (
	"fmt"
	"path/filepath"
	"os/exec"
	"os"
	"github.com/sas/utils"
)

func main() {
	testtargzip()
}

func testtargzip() {
	path := "/home/juken/gopath/src/github.com/sas/sasslvpn_web/modules"
	file := "/home/juken/gopath/src/github.com/sas/sasslvpn_web/modules/targz.tar.gz"
	err := CompressTarGzipByCommand(path + "/zgf", file)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DeCompressTarGzipByCommand(file, path + "/targz")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CompressTarGzipByCommand(path, fileName string) error {
	path = filepath.Clean(path)
	cmd := exec.Command("tar", "zcvf", fileName, filepath.Base(path))
	cmd.Dir = filepath.Dir(path)
	return cmd.Run()
}

func DeCompressTarGzipByCommand(fileName, path string) error {
	path = filepath.Clean(path)
	//删除目录及以下文件
	if utils.CheckFileExist(path) {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	//创建一个目录
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	//解压
	cmd := exec.Command("tar", "zxvf", fileName)
	cmd.Dir = path
	return cmd.Run()
}