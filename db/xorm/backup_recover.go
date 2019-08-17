package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//backup()
	recover()
}

func backup() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "localhost:3306", "gkzx") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	err = Engine.DumpAllToFile("d://test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func recover() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "localhost:3306", "gkzx") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	tables, err := Engine.DBMetas()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range tables {
		res,err:=Engine.Exec("drop table "+v.Name)
		fmt.Println(res,err,"..............")
	}
	_, err = Engine.ImportFile("d://test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
}