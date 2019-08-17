package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type SzgaUserHeartbeatHis struct {
	Id                  int64    `json:"id"`
	TerminalHwId        string    `json:"terminal_hw_id"`    //终端硬件ID
	PoliceId            string    `json:"police_id"`         //警号
	HeartbeatTime       time.Time `json:"heartbeat_time"`    //终端心跳时间
	Online              bool      `json:"online"`            //TRUE上线，FALSE下线
	OnlineTime          time.Time `json:"online_time"`       //上线时间
	OfflineTime         time.Time `json:"offline_time"`      //下线时间
	StateChangeTime     time.Time `json:"state_change_time"` //状态改变时间
	TerminalId          string    `json:"terminal_id"`       //终端ID
	Imei                string    `json:"imei"`              //
	Imsi                string    `json:"imsi"`              //
	Mac                 string    `json:"mac"`               //
	ApnvpdnList         string    `json:"apnvpdn_list"`      //
	CertSn              string    `json:"sert_sn"`           //证书sn
	Longitude           float64   `json:"longitude"`         //经度
	Latitude            float64   `json:"latitude"`          //纬度

	AddressAdcode       string `json:"address_adcode"`
	AddressCity         string `json:"address_city"`
	AddressCountry      string `json:"address_country"`
	AddressCountryCode  string `json:"address_country_code"`
	AddressDirection    string `json:"address_direction"`
	AddressDistance     string `json:"address_distance"`
	AddressDistrict     string `json:"address_district"`
	AddressProvince     string `json:"address_province"`
	AddressStreet       string `json:"address_street"`
	AddressStreetNumber string `json:"address_street_number"`

	UpdatedAt           time.Time `json:"updated_at"xorm:"updated"`
	ChangeStatus        bool     `json:"change_status"`      //是否改变上下线状态
}

func (s SzgaUserHeartbeatHis) TableName() string {
	return "szga_user_heartbeat_his"
}

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "192.168.128.159:3306", "test") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(false)
	err = Engine.Sync2(
		new(SzgaUserHeartbeatHis),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	his := new(SzgaUserHeartbeatHis)
	his.ChangeStatus = true
	/////////////////////
	his.TerminalHwId = "TerminalHwId"
	his.PoliceId = "PoliceId"
	his.HeartbeatTime = time.Now()
	his.Online = true
	his.OnlineTime = time.Now()
	his.OfflineTime = time.Now()
	his.StateChangeTime = time.Now()
	his.TerminalId = "TerminalId"
	his.Imei = "Imei"
	his.Imsi = "Imsi"
	his.Mac = "Mac"
	his.ApnvpdnList = "ApnvpdnList"
	his.CertSn = "CertSn"
	his.Longitude = 103.3242
	his.Latitude = 34.23543564
	his.AddressAdcode = "645874"
	his.AddressCity = "510521"
	his.AddressCountry = "510521"
	his.AddressCountryCode = "510521"
	his.AddressDirection = "AddressDirection"
	his.AddressDistance = "2432"
	his.AddressDistrict = "AddressDistrict"
	his.AddressProvince = "AddressProvince"
	his.AddressStreet = "AddressStreet"
	his.AddressStreetNumber = "AddressStreetNumber"
	start := time.Now().Unix()
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 5000; i++ {
		//fmt.Println(i)
		his.Id = 0
		//_, err = Engine.Insert(his)
		_, err = session.Insert(his)
		if err != nil {
			fmt.Println(err, i)
			break
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now().Unix()
	fmt.Println("用时：", end - start, "秒")
}
