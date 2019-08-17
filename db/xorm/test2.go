package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sas/tqry/app/models"
)

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "localhost:3306", "tqry") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	err = Engine.Sync2(
		new(models.SysConfig),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	var curogr = models.SysConfig{ConfigKey:models.PUBLISH_ORG, ConfigValue:"fdsaf", ConfigDesc:"2222"}
	var conf models.SysConfig
	exist, err := Engine.Id(curogr.ConfigKey).Get(&conf)
	fmt.Println(exist, err,conf)
	if conf.ConfigKey == "" {
		_, err = Engine.Insert(curogr)
	} else {
		_, err = Engine.Id(curogr.ConfigKey).Update(curogr)
	}
	fmt.Println(err)
}
