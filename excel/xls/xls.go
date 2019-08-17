package main

import (
	"fmt"
	"github.com/extrame/xls"
)

func main() {
	xlFile, err := xls.Open("E:/GOPATH/src/testgo/excel/xls/event_type.xls", "utf-8")
	if err != nil {
		fmt.Println(err)
		return
	}
	if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
		total := int(sheet1.MaxRow)
		for i := 0; i <= total; i++ {
			row := sheet1.Row(i)
			if row == nil {
				continue
			}
			fmt.Println(row.Col(0))
			fmt.Println(row.Col(1))
		}
	}
}
