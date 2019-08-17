package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sas/gkzx/app/models"
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/sas/utils"
)

func main() {
	syncUser()
	//syncOrg()
	//app()
}

func getEngine() (*xorm.Engine, error) {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "192.168.130.240:3306", "gkzx") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return Engine, nil
}

func syncUser() error {
	engine, err := getEngine()
	if err != nil {
		return err
	}
	var users[] models.UsersInfo
	err = engine.Where("deleted=?", 0).Find(&users)
	if err != nil {
		fmt.Println(err)
		return err
	}
	jsonArray := []*simplejson.Json{}
	for _, v := range users {
		jsonitem := simplejson.New()
		jsonitem.Set("__time", time.Now().Unix() * 1000)
		jsonitem.Set("user_id", v.UserId)
		jsonitem.Set("dept_id", v.UserOrg)
		jsonitem.Set("mobile", v.Telephone)
		jsonitem.Set("name", v.UserName)
		jsonitem.Set("police_id", v.PoliceId)
		jsonitem.Set("police_type", v.PoliceType)
		jsonitem.Set("user_ic", v.UserId)
		jsonitem.Set("register_time", v.RegTime.Unix() * 1000)
		jsonitem.Set("modify_time", v.RegTime.Unix() * 1000)
		jsonitem.Set("is_delete", v.Deleted)
		jsonitem.Set("report_time", time.Now().Unix() * 1000)
		jsonArray = append(jsonArray, jsonitem)
	}
	return utils.IndexingBranch("https://192.168.130.240:9007/search/bulk", "szga", "szga_users_info", jsonArray)
}

func syncOrg() error {
	engine, err := getEngine()
	if err != nil {
		return err
	}
	var orgs[] models.Org
	err = engine.Find(&orgs)
	if err != nil {
		fmt.Println(err)
		return err
	}
	jsonArray := []*simplejson.Json{}
	for _, v := range orgs {
		jsonitem := simplejson.New()
		jsonitem.Set("__time", time.Now().Unix() * 1000)
		jsonitem.Set("dept_id", v.Id)
		jsonitem.Set("dept_name", v.Name)
		jsonitem.Set("pid", v.Parent)
		jsonitem.Set("is_delete", 0)
		jsonitem.Set("register_time", v.CreatedAt.Unix() * 1000)
		jsonitem.Set("modify_time", v.UpdatedAt.Unix() * 1000)
		jsonitem.Set("report_time", time.Now().Unix() * 1000)
		jsonArray = append(jsonArray, jsonitem)
	}
	return utils.IndexingBranch("https://192.168.130.240:9007/search/bulk", "szga", "szga_sys_organization", jsonArray)
}
//应用数据
func app() {
	jsonArray := []*simplejson.Json{}
	m := map[string]string{"qq":"QQ", "com.anrong.weixin":"微信"}
	for k, v := range m {
		jsonitem := simplejson.New()
		jsonitem.Set("__time", time.Now().Unix() * 1000)
		jsonitem.Set("app_code", k)
		jsonitem.Set("app_name", v)
		jsonitem.Set("app_type", "应用类别")
		jsonitem.Set("app_version", "版本号")
		jsonitem.Set("build_dept_id", "建设单位")
		jsonitem.Set("custodian_email", "zrryx@163.com")
		jsonitem.Set("custodian_id", "责任人")
		jsonitem.Set("custodian_telephone", "13500000009")
		jsonitem.Set("domain_name", "wondersoft")
		jsonitem.Set("is_delete", 0)
		jsonitem.Set("maintenance_dept_id", "维护单位")
		jsonitem.Set("maintenance_man_email", "wfwyx@163.com")
		jsonitem.Set("maintenance_man_id", "维护人")
		jsonitem.Set("maintenance_man_telephone", "13500000009")
		jsonitem.Set("modify_time", 1483203670000)
		jsonitem.Set("register_time", time.Now().Unix() * 1000)
		jsonitem.Set("server_ip", "127.0.0.1")
		jsonitem.Set("server_ports", "9999")
		jsonArray = append(jsonArray, jsonitem)
	}
	utils.IndexingBranch("https://192.168.130.240:9007/search/bulk", "szga", "szga_app_info", jsonArray)
}