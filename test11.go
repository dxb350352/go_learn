package main
import (
	"github.com/sas/utils"
	"github.com/bitly/go-simplejson"
	"fmt"
	"time"
)


func main() {
	url := "http://192.168.130.240:9005/search/indexing"
	__app := "szga"
	__type := "szga_users_info"
	var data []simplejson.Json
	js := simplejson.New()
	js.Set("id", 1231)
	js.Set("__time", time.Now().Unix())
	data = append(data, *js)
	err := utils.IndexingBranch(url, __app, __type, data)
	fmt.Println(err)
	fmt.Println(utils.GetDirNameForGZIP("C:/Users/Administrator/Desktop/wsv3.tar.gz"))
}