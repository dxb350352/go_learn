package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/ss.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var provice []*simplejson.Json
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			name, _ := row.Cells[1].String()
			code, _ := row.Cells[0].String()
			pcode, _ := row.Cells[2].String()
			area := simplejson.New()
			area.Set("id", code)
			area.Set("name", name)
			area.Set("pId", pcode)
			provice = append(provice, area)
		}
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
