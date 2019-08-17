package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sas/gkzx/app/models"
	"strings"
	"github.com/sas/utils"
	"os"
	"io/ioutil"
	"fmt"
)

var clsFolderPath = "d:/menu/cls/"
var htmlFolderPath = "d:/menu/html/"
var htmltemplate = "d:/userinfoReports.html"

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.240:3306", "gkzx") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	var menus []models.Menu
	err = Engine.Where("enable = ?", 1).OrderBy("url").Find(&menus)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !utils.CheckFileExist(clsFolderPath) {
		os.MkdirAll(clsFolderPath, 0775 | os.ModeDir)
	}
	if !utils.CheckFileExist(htmlFolderPath) {
		os.MkdirAll(htmlFolderPath, 0775 | os.ModeDir)
	}
	var s string
	var clsname string
	var arr[]string
	for _, v := range menus {
		arr = strings.Split(v.Url, "/")
		index := strings.Index(arr[2], "?")
		if index != -1 {
			arr = append(arr[:2], arr[2][:index])
		}
		//后台
		if len(arr) == 3 && arr[2] != "" {
			//另一个类
			if clsname != "" &&clsname != arr[1] {
				tofile(s, clsname)
				s = ""
			}
			s += `
			func (this *` + upper(arr[1]) + `) ` + upper(arr[2]) + `() revel.Result {
				return this.Render()
			}
			`
			clsname = arr[1]
		}
		//html
		p := htmlFolderPath + arr[1] + "/"
		if !utils.CheckFileExist(p) {
			os.MkdirAll(p, 0775 | os.ModeDir)
		}
		utils.CopyFile(p + arr[2] + ".html", htmltemplate)
	}
	tofile(s, arr[1])
}

func upper(str string) string {
	var byt[] byte
	byt = append(byt, str[0] - 32)
	byt = append(byt, str[1:]...)
	return string(byt)
}

func tofile(content, filename string) {
	fname := clsFolderPath + filename + ".go"
	fc := []byte(content)
	if utils.CheckFileExist(fname) {
		fct, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println(err)
			return
		}
		fc = append(fct, fc...)
	} else {
		content = `package controllers

					import (
						"github.com/revel/revel"
					)

					type ` + upper(filename) + ` struct {
						*revel.Controller
					}` + content
		fc = []byte(content)
	}
	ioutil.WriteFile(fname, fc, os.ModePerm)
}