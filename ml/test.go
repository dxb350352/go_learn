package main

import (
	"fmt"
	"time"
	"github.com/bitly/go-simplejson"
	"github.com/jackiedong168/gorequest"
	"errors"
)

func main() {

}

func mlcheck(j *simplejson.Json) error {
	j.Set("is_alarm", "0")
	type1 := j.Get("dwTypeId1").MustString()
	type2 := j.Get("dwTypeId2").MustString()
	//只验证U盘插拔
	if type1 == "28672" && type2 == "28680" {
		__time, err := j.Get("__time").Int64()
		if err != nil {
			return err
		}
		data := fmt.Sprintf("%s,%s,%d", type1, type2, (__time / 1000 ) % 86400)
		__time = __time / 1000
		__time = __time % 86400
		fmt.Println(data, __time)

		goReq := gorequest.New()

		rr := goReq.Timeout(time.Minute).Get("http://192.168.130.240:8088/sparkml/mlcheck").Param("data", data)

		_, body, errs := rr.EndBytes()
		if len(errs) != 0 {
			return errors.New("查找数据出错")
		}
		bodyjs, err := simplejson.NewJson(body)
		if nil != err {
			return err
		}
		if bodyjs.Get("status").MustInt64() == 200 {
			result := bodyjs.Get("data").MustFloat64()
			if result == 1.0 {
				j.Set("is_alarm", "1")
			} else {
				j.Set("is_alarm", "0")
			}
		}
	}
	return nil
}