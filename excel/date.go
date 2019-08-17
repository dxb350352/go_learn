package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/User.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			result, _ := strconv.ParseInt(row.Cells[3].Value, 10, 64)
			if result > 0 {
				fmt.Println(time.Unix((result-25569)*24*60*60, 0).Format("2006-01-02"))
			}
		}
	}
}
