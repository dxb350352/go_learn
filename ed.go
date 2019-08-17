package main
import (
	"time"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	aa()
}

func gg() {
	ddate, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
	}
	ddate = ddate.AddDate(0, 0, 1 - 30)
	fmt.Println(ddate)
}

func aa() {
	//删除数据库
	ddate, err := time.Parse("2006-01-02 15:04:05", "2016-04-26 11:00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ddate.Unix() * 1000)
	//删除search
	///search/deletealldata (app, _type, interval string, start_time, end_time int64)
	url := "http://192.168.130.240:9006/search/deletealldata"
	request := gorequest.New()
	_, body, errs := request.Get(url).Query("app=szga_monitorsvr").Query("_type=t_termonline").Query(fmt.Sprint("start_time=", 0)).Query(fmt.Sprint("end_time=", ddate.Unix() * 1000)).End()
	if errs != nil {
		fmt.Println(err)
		return
	}
	j2, err := simplejson.NewJson([]byte(body))
	fmt.Println(j2)
	if err != nil {
		fmt.Println(err)
		return
	}
	if j2.Get("status").MustInt() != 200 {
		fmt.Println(err)
	}
}
