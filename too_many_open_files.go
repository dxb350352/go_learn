package main
import (
	"github.com/jackiedong168/gorequest"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	for i := 0; i < 2000; i++ {
		test(i)
	}
}

func test(i int) {

	//修改gorequest,main.go---594行，req.Close=true
	request := gorequest.New()
	_, body, errs := request.Get("http://192.168.130.240:9006/search/search").Query(fmt.Sprint("keywords=__app=szga_monitorsvr AND __type=t_termonline AND PoliceID=", 4)).Query(fmt.Sprint("startTime=", 0)).Query(fmt.Sprint("endTime=", 0)).End()
	if errs != nil {
		fmt.Println(errs)
		return
	}
	_, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println("返回数据不是JSON" + body)
		return
	}
	fmt.Println(i)
}