package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "localhost:3306", "tqry") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		log.Fatal(err)
	}
	list, err := Engine.SQL("select informer_no from informers").QueryString()
	if err != nil {
		log.Fatal(err)
	}
	for i, m := range list {
		fmt.Println(i, m["informer_no"][:12])
	}
	time.Sleep(time.Hour)
}
