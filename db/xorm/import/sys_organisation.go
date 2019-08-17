package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
	"time"
)

type SysOrganisation struct {
	Dwdm   string `json:"dwdm" xorm:"pk VARCHAR(12)" colName:"单位代码"`    //单位代码
	Dwmc   string `json:"dwmc" xorm:"VARCHAR(200)" colName:"单位名称"`      //单位名称
	Sjdm   string `json:"sjdm" xorm:"VARCHAR(12) index" colName:"上级单位"` //上级单位
	Dwjc   string `json:"dwjc" xorm:"VARCHAR(200)" colName:"单位简称"`      //单位简称
	Yxx    string `json:"yxx" xorm:"VARCHAR(1)" colName:"有效性"`          //有效性
	Dwjb   string `json:"dwjb" xorm:"VARCHAR(1)" colName:"单位级别"`        //单位级别
	Xzqh   string `json:"xzqh" xorm:"VARCHAR(6)" colName:"所属地区"`        //所属地区
	Dwjz   string `json:"dwjz" xorm:"VARCHAR(4)" colName:"单位警种"`        //单位警种
	Ywsjdm string `json:"ywsjdm" xorm:"VARCHAR(12)" colName:"业务上级单位"`   //业务上级单位
}

func (s *SysOrganisation) TableName() string {
	return "sys_organisations_2"
}

var dataReg = regexp.MustCompile(`values \((.*)\);`)

var total int
var success int
var remaindata = []string{
	"21",
	"3",
	"42",
	"5",
	"804",
	"901",
	"903",
	"904",
}

func main() {
	start := time.Now().Unix()
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.201:3306", "tqry") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open("E:/GOPATH/src/testgo/db/xorm/import/dwdm.sql")
	scan := bufio.NewScanner(file)
	var dbs []*SysOrganisation
	var line string
	for scan.Scan() {
		line += strings.TrimSpace(scan.Text())
		if strings.HasSuffix(line, ";") {
			aa := dataReg.FindAllStringSubmatch(line, -1)
			if len(aa) > 0 && len(aa[0]) > 1 {
				line = aa[0][1]
				line = strings.Replace(line, "'", "", -1)
				arr := strings.Split(line, ",")
				if len(arr) >= 9 {
					//去掉不要的警种
					jz := strings.TrimLeft(strings.TrimSpace(arr[7]), "0")
					if !ContainsString(remaindata, jz) {
						continue
					}
					dbs = append(dbs, &SysOrganisation{Dwdm: arr[0],
						Dwmc: strings.TrimSpace(arr[1]),
						Sjdm: strings.TrimSpace(arr[2]),
						Dwjc: strings.TrimSpace(arr[3]),
						Yxx: strings.TrimSpace(arr[4]),
						Dwjb: strings.TrimSpace(arr[5]),
						Xzqh: strings.TrimSpace(arr[6]),
						Dwjz: jz,
						Ywsjdm: strings.TrimSpace(arr[8])})
					if len(dbs) >= 200 {
						BatchInsert(Engine, dbs)
						dbs = []*SysOrganisation{}
					}
				}
			}
			line = ""
		}
	}
	if len(dbs) > 0 {
		BatchInsert(Engine, dbs)
	}
	fmt.Println("total:", total, "success:", success, "takes time:", time.Now().Unix()-start)
}

func Batch(Engine *xorm.Engine, datas []*SysOrganisation) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.Insert(datas)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}

func BatchInsert(Engine *xorm.Engine, datas []*SysOrganisation) {
	total += len(datas)
	var failure int
	err := Batch(Engine, datas)
	if err != nil {
		for i, _ := range datas {
			_, err := Engine.InsertOne(datas[i])
			if err != nil {
				failure++
				fmt.Println(err)
			}
		}
	}
	success += len(datas) - failure
}

func ContainsString(arr []string, target string) bool {
	for _, o := range arr {
		if o == target {
			return true
		}
	}
	return false
}
