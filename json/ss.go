package main
import (
	"fmt"
	js "github.com/bitly/go-simplejson"
)

func main() {
	str := "{\"%msgtime% [ err ] e : /gopath_go/src/github.com/sas/modules/auth/app/models/init.go err : %string% { op : ' %action% ' , path : ' %string% ' , err : 0x2 }\":{\"ex\":\"2015-09-06/18:04:55 [ERR] e:/GOPATH_GO/src/github.com/sas/modules/auth/app/models/init.go err:&os.PathError{Op:'open', Path:'./sas.rn', Err: 0x2}\",\"cnt\":54},\"%msgtime% [ err ] e : /gopath_go/src/github.com/sas/modules/auth/app/models/init.go err : %string% { s : ' %string% ' }\":{\"ex\":\"2015-09-06/18:04:55 [ERR] e:/GOPATH_GO/src/github.com/sas/modules/auth/app/models/init.go err:&errors.errorString{s:'获取硬件ID失败'}\",\"cnt\":27}}"
	//	fmt.Println(str)
	m1, err := js.NewJson([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
	m := m1.MustMap()
	for k, v := range m {
		vv := v.(map[string]interface{})
		fmt.Println(k)
		fmt.Println(vv["ex"])
		fmt.Println(vv["cnt"])
	}
	m1.Set("ttt",[]string{"1","2"})
	fmt.Println(m1.Get("ttt").MustString(),"............")
}