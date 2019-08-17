package main

import (
	"github.com/bitly/go-simplejson"
	"log"
	"fmt"
)

func main() {
	str := `{
    "__app": "szga",
    "__index": "sas_szga_szga_users_info",
    "__time": 1501032338000,
    "__type": "szga_users_info",
    "dept_id": "640521210000",
    "is_delete": 0,
    "mobile": "13689926770",
    "modify_time": 1501032338000,
    "name": "走一个",
    "police_id": "211086",
    "police_type": "010",
    "register_time": 1501032338000,
    "report_time": 1501036873000,
    "user_id": "2b67a7db2bc5baaffabeafb2a4cc8ef9"
}`
	data, err := simplejson.NewJson([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
	j := simplejson.New()
	j.Set("__app", data.Get("__app").MustString())
	j.Set("__type", data.Get("__type").MustString())
	j.Set("data", []*simplejson.Json{data})
	byt, _ := j.MarshalJSON()
	fmt.Printf(`curl -l -k -H "Content-type: application/json" -X POST -d '%s' https://localhost:9005/search/bulk?by=all`, byt)

}

