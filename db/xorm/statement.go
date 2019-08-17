package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.201:3306", "gaxz") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = Engine.Exec(`INSERT INTO score_1 (id,score) VALUES (?,?),(?,?),(?,?)`, "1", "1","2","2","3","3")
	fmt.Println(err)
}
