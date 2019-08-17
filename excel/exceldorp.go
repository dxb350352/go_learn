package main
import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"github.com/bitly/go-simplejson"
	"encoding/json"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/ss.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var provice []*simplejson.Json
	var city []*simplejson.Json
	var district []*simplejson.Json
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			name, _ := row.Cells[1].String()
			code, _ := row.Cells[0].String()
			pcode, _ := row.Cells[2].String()
			level, _ := row.Cells[4].String()
			area := simplejson.New()
			area.Set("code", code)
			area.Set("name", name)
			area.Set("pcode", pcode)
			if level == `1` {
				provice = append(provice, area)
			}else if level == `2` {
				city = append(city, area)
			}else if level == `3` {
				district = append(district, area)
			}
		}
	}
	for i, v := range provice {
		var children []*simplejson.Json
		for _, vv := range city {
			if v.Get("code").MustString() == vv.Get("pcode").MustString() {
				var dchildren []*simplejson.Json
				for _, vvv := range district {
					if vv.Get("code").MustString() == vvv.Get("pcode").MustString() {
						dchildren = append(dchildren, vvv)
					}
				}
				vv.Set("children", dchildren)
				children = append(children, vv)
			}
		}
		provice[i].Set("children", children)
	}

	file, err := os.Create(`E:/GOPATH_GO/src/testgo/excel/ss.json`)
	if err != nil {
		fmt.Println(err)
		return
	}
	temp, err := json.Marshal(provice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(temp))
	_, err = file.Write(temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}