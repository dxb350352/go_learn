package main

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sas/utils"
	"github.com/bitly/go-simplejson"
	"github.com/sas/gkzx/app/models"
	sttmod "github.com/sas/modules/strategy/app/models"
	"encoding/json"
)
//离线长--单位分钟
var offline_minute int64 = 30
//上线后一段时间--单位分钟
var online_handle_minute int64 = 30000

func main() {
	//需要两个参数加上默认有一个参数，一共至少3个
	if len(os.Args) < 3 {
		fmt.Print("[]")
		return
	}
	resultUrl := sttmod.BuildUrl(os.Args[1])
	dataSourceName := os.Args[2]
	//获取一段时间内终端访问应用次数
	request, err := utils.GetSkipValidateRequest(resultUrl)
	if err != nil {
		fmt.Print("[]")
		return
	}
	//查询
	_, body, errs := request.Get(resultUrl).EndBytes()
	if len(errs) > 0 {
		fmt.Print("[]")
		return
	}
	j, err := simplejson.NewJson(body)
	length := len(j.MustArray())
	if length == 0 {
		fmt.Print("[]")
		return
	}
	//获取离线时间超过一定时间的终端并刚上线不久
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Print("[]")
		return
	}
	sql := `SELECT
			*
		FROM
			szga_terminal_heartbeat t
		WHERE
			t. ONLINE = TRUE
		AND TIMESTAMPDIFF(
			MINUTE,
			t.offline_time,
			t.online_time
		) > ?
		AND TIMESTAMPDIFF(MINUTE, t.online_time, NOW()) < ?
		AND t.terminal_hw_id in(`
	for i := 0; i < length; i++ {
		t := j.GetIndex(i).Get("key").MustString()
		sql += `"` + t + `",`
	}
	sql = sql[:len(sql) - 1]
	sql += ")"
	var dbs[] models.SzgaTerminalHeartbeat
	err = engine.Sql(sql, offline_minute, online_handle_minute).Find(&dbs)
	if err != nil {
		fmt.Print("[]")
		return
	}
	var result []*simplejson.Json
	for _, v := range dbs {
		for i := 0; i < length; i++ {
			if v.TerminalHwId == j.GetIndex(i).Get("key").MustString() {
				byt, _ := json.Marshal(v)
				re, _ := simplejson.NewJson(byt)
				re.Set("count", j.GetIndex(i).Get("doc_count"))
				result = append(result, re)
				break
			}
		}
	}
	byt, _ := json.Marshal(result)
	fmt.Print(string(byt))
}

