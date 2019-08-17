package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//	file, err := os.Open("./src/testgo/xorm/auth_menus.sql")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "192.168.128.159:3306", "gkzx") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	path := "E:/GOPATH/src/testgo/db/xorm"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		file, err := os.Open(path + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.HasSuffix(f.Name(), ".sql") {
			fmt.Println(f.Name())
			Engine.Import(file)
		}
	}
}
