package main
import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	ss := `{"a":1,"b":2.0,"c":3.3}`
	json, err := simplejson.NewJson([]byte(ss))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(json.Get("a").MustInt64())
	fmt.Println(json.Get("b").MustFloat64())
	fmt.Println(json.Get("c").MustFloat64())
	fmt.Println(json.Get("c").MustInt64())
	fmt.Println(json.Get("a").MustFloat64())
	bb, _ := json.MarshalJSON()
	fmt.Println(string(bb))
}
