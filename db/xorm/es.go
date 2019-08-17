package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/json"
	"github.com/jackiedong168/gorequest"
)

type Person struct {
	Id            int64  `json:"_id"`
	PersonName    string `json:"person_name"`
	PersonPhone   string `json:"person_phone"`
	PersonId      string `json:"person_id"`
	PersonQq      string `json:"person_qq"`
	PersonEmail   string `json:"person_email"`
	PersonAddress string `json:"person_address"`
	PersonGender  string `json:"person_gender"`
}

func (p Person) TableName() string {
	return "person_info"
}

var esurl = "http://192.168.130.201:19200/nbo_myperson/myperson"

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.201:3306", "person") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		log.Fatal(err)
	}
	err = Engine.Sync2(
		new(Person),
	)
	if err != nil {
		log.Fatal(err)
	}
	var dbs []Person
	err = Engine.Find(&dbs)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dbs {
		byt, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		request := gorequest.New()
		response, body, errs := request.Put(fmt.Sprintf("%s/%d", esurl, v.Id)).Send(byt).End()
		fmt.Println(response,body,errs)
	}

}
