package main

import (
	"fmt"
	"os"
	"log"
	"math/rand"
	"time"
)

//新府新谷6号楼
var tfxg_lat = 30.592861
var tfxg_lon = 104.061523
//仁和春天5幢
var rhct_lat = 30.592927
var rhct_lon = 104.058536
//坐标最大变化范围
var gap_max = 0.0003
var header1 = `person_weixin|person_nick|person_location_gps|weixin_title|weixin_content|time_happen|person_location_desc|person_image`
var header2 = `person_id|person_weixin|person_nick|person_location_gps|weixin_title|weixin_content|time_happen|person_location_desc|person_image`

func main() {
	tfxg, err := os.Create("E:/GOPATH/src/testgo/666-111-1-1-100.txt", )
	if err != nil {
		log.Fatal(err)
	}
	defer tfxg.Close()
	tfxg.WriteString(header1)
	rhct, err := os.Create("E:/GOPATH/src/testgo/666-222-1-1-100.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer rhct.Close()
	rhct.WriteString(header2)

	weixin_start := 654321
	person_id_start := "51010120180101"
	person_location_desc:="四川省成都市武侯区府城大道888号"
	person_image:="https://www.elastic.co/assets/blt42262f54b685e96a/icon-white-circle-elastic-search.svg"
	rand.Seed(time.Now().UnixNano())
	timestr := time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < 100; i++ {
		tfxg.WriteString(fmt.Sprintf("\n%d|昵称%d|%s|标题%d|内容%d|%s|%s|%s", weixin_start+i, i, get_random_coor(tfxg_lat, tfxg_lon), i, i, timestr,person_location_desc,person_image))
		rhct.WriteString(fmt.Sprintf("\n%s%04d|%d|nick%d|%s|title%d|content%d|%s|%s|%s", person_id_start, i, weixin_start+i, i, get_random_coor(rhct_lat, rhct_lon), i, i, timestr,person_location_desc,person_image))
	}
}

func get_random_coor(lat, lon float64) string {
	tgap := int(gap_max * 1000000)
	return fmt.Sprintf("%f,%f", lat+gap_max-float64(rand.Intn(tgap))/1000000.0, lon+gap_max-float64(rand.Intn(tgap))/1000000.0, )
}
