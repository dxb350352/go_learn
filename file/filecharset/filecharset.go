package main
import (
	"io/ioutil"
	"github.com/saintfish/chardet"
	"fmt"
	"errors"
)

func main() {
	filepath := "E:/GOPATH_GO/src/testgo/file/filecharset/other"
	cs := getcharset(filepath)
	fmt.Println(cs)
	if cs == "" {
		errors.New("请上传utf-8格式的文件")
	}else if cs == "UTF-8" {

	}else {
//		filebyte, err := ioutil.ReadFile(filepath)
//		if err != nil {
//			fmt.Println(err)
//		}
	}
}

func getcharset(filepath string) string {
	fb, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result, err := chardet.NewTextDetector().DetectBest(fb)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result.Charset
}
