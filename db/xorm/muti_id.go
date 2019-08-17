package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type AuthMenu struct {
	Province string `xorm:"pk"`
	City     string `xorm:"pk"`
	UserName string `xorm:"pk"`
	Menu     string
}

func (this AuthMenu) TableName() string {
	return "auth_menu"
}

var Engine *xorm.Engine

func (this *AuthMenu) Get(province, city, userName string) error {
	_, err := Engine.Where("province=?", province).And("city=?", city).And("user_name=?", userName).Get(this)
	return err
}

func main() {
	var err error
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.201:3306", "gaxz") + "&loc=Asia%2FChongqing"
	Engine, err = xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)

	var auth AuthMenu
	fmt.Println(auth.Get("全国", "全省", "测试"), auth)
	fmt.Println(auth)
}
