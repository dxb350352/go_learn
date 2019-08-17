package main
import (
	"io/ioutil"
	"os"
	"regexp"
	"fmt"
	"errors"
)

func main() {
//		fmt.Println(AddRoutes("v3"))
	fmt.Println(RemoveRoutes("v3"))
}

func AddRoutes(moduleName string) error {
	//级数最终添加内容
	module := fmt.Sprintln("GET", "/public/" + moduleName + "/*filepath", `Static.ServeModule("` + moduleName + `","public/` + moduleName + `")`)
	//读取原文件
	content, err := ioutil.ReadFile("E:/GOPATH_GO/src/testgo/regexp/routes")
	if err != nil {
		return err
	}
	//判断是否存在
	reg := regexp.MustCompile(`((GET)|(POST))\s+(/public/` + moduleName + `/\*filepath)\s+(Static.ServeModule\("` + moduleName + `","public/` + moduleName + `"\))\s*`)
	if reg.Match(content) {
		return errors.New("已经存在")
	}
	//准备写入文件
	f, err := os.Create("E:/GOPATH_GO/src/testgo/regexp/routes")
	if err != nil {
		return err
	}
	defer f.Close()
	reg = regexp.MustCompile(`((GET)|(POST))\s+(/public/\*filepath)\s+(Static.Serve\("public"\))`)

	result := reg.FindAll(content, 1)
	for _, v := range result {
		//把module添加在reg之前
		content = reg.ReplaceAll(content, append([]byte(module), v...))
	}
	_, err = f.Write(content)
	return err
}

func RemoveRoutes(moduleName string) error {
	//读取原文件
	content, err := ioutil.ReadFile("E:/GOPATH_GO/src/testgo/regexp/routes")
	if err != nil {
		return err
	}
	reg := regexp.MustCompile(`((GET)|(POST))\s+(/public/` + moduleName + `/\*filepath)\s+(Static.ServeModule\("` + moduleName + `","public/` + moduleName + `"\))\s*`)
	if reg.Match(content) {
		content = reg.ReplaceAll(content, []byte{})
		//准备写入文件
		f, err := os.Create("E:/GOPATH_GO/src/testgo/regexp/routes")
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(content)
	}
	return err
}
