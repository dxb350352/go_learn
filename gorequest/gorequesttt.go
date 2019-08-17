package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
	"strconv"
	"github.com/hpcloud/tail"
	"os"
	"github.com/bitly/go-simplejson"
	"time"
	"github.com/robfig/config"
)

type PathConfInfo struct {
	dir           string
	filename      string
	posturl       string
	LogApp        string
	LogType       string

	FindFileSpace time.Duration
	////////////////////////
	//Request *gorequest.SuperAgent
}

const REC_CONFIG string = "./rec_conf/"

func main() {
	//for i := 0; i < 20000; i++ {
	//	var RequestGo = goreq.New()
	//	_, body, errs := RequestGo.Get("http://192.168.130.181:8080/tm/query?dsn=jdbc:informix-sqli://192.168.130.160:9090/db_test:&driver=com.informix.jdbc.IfxDriver&user=AAO&password=123456&sql=select%20*%20from%20systables").End()
	//	fmt.Println(i, "-------", string(body), errs)
	//}
	pathfile := "E:/GOPATH/src/github.com/sas/sasclient/common.go"
	var conf PathConfInfo
	conf.LogApp = "xy"
	conf.LogType = "xy"
	conf.posturl = "http://192.168.130.121:9005/DataStream/up"
	conf.FindFileSpace = time.Second * 10
	FileLogup(pathfile, &conf, nil)
}

func FileLogup(pathfile string, conf *PathConfInfo, r *gorequest.SuperAgent) {
	pathfile2 := strings.Replace(pathfile, "/", "_", -1)
	pathfile2 = strings.Replace(pathfile2, ".", "_", -1)
	//读取配置
	path := REC_CONFIG + pathfile2
	cfg, _ := config.Read(path, config.DEFAULT_COMMENT, config.ALTERNATIVE_SEPARATOR, false, true)
	if nil == cfg {
		cfg = config.New(config.DEFAULT_COMMENT, config.ALTERNATIVE_SEPARATOR, false, true)
	}
	//record line count
	iLine, _ := cfg.Int("config", "line_cnt")
	OffSet, _ := cfg.String("config", "offset")
	iOffSet, _ := strconv.ParseInt(OffSet, 10, 64)

	tC := tail.Config{Follow: true}
	tC.Location = &tail.SeekInfo{iOffSet, os.SEEK_SET}
	t, err := tail.TailFile(pathfile, tC)
	if nil != err {
		fmt.Println("TailFile ", pathfile, ", err:", err)
		return
	}

	if nil == r {
		r = gorequest.New()
	}

	for line := range t.Lines {
		if nil != line.Err {
			fmt.Println("tail -f err:", line.Err)
			break
		}
		strLen := len(line.Text)
		if strLen == 0 {
			break
		}
		iOffSet += int64(strLen) + 1
		_, body, errs := r.Post(conf.posturl).Param("__app", conf.LogApp).Param("__type", conf.LogType).
		Param("__id", pathfile2 + "_" + strconv.FormatInt(int64(iLine), 10)).Type("text").SendString(line.Text).EndBytes()
		iLine++
		if nil != errs {
			fmt.Println("request error:", errs)
			break
		}
		js, err := simplejson.NewJson(body)
		if nil != err || js.Get("status").MustInt() != 200 {
			fmt.Println("response body:", string(body))
			break
		} else {
			cfg.AddOption("config", "offset", strconv.FormatInt(iOffSet, 10))
			cfg.AddOption("config", "line_cnt", strconv.Itoa(iLine))
			cfg.WriteFile(path, os.ModePerm, "line_cnt为已经抓取了多少条日志")
			fmt.Println("path:", pathfile, ", line:", iLine)
		}
		//time.Sleep(conf.FindFileSpace)
	}
	time.Sleep(conf.FindFileSpace)
	go FileLogup(pathfile, conf, r)
}