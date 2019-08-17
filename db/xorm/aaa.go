package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type IndexAlias struct {
	Id string `xorm:"index pk"`
}

func (this IndexAlias) TableName() string {
	return "gaxz_index_alias"
}
func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.201:3306", "gaxz") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Engine.Sync2(
		new(IndexAlias),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	var db IndexAlias
	exist,err:=Engine.Id("aliasgax1z_object_groupdesc").Get(&db)
	fmt.Println(exist,err,db)
}
