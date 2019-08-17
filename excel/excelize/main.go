package main

import (
	"github.com/Luxurioust/excelize"
	"fmt"
)

func main() {
	test()
}
func test() {
	xlFile, err := excelize.OpenFile("E:/GOPATH/src/testgo/excel/statis_model.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	xlFile.SetCellValue("Sheet1","A2","填报单位：fdsafdsaf")
	xlFile.SetCellValue("Sheet1","V2","统计期间：2016年12月12日至2016年12月21日 ")
	xlFile.SetCellValue("Sheet1","A6","321")
	xlFile.SetCellValue("Sheet1","B6","24")
	xlFile.SetCellValue("Sheet1","C6","24")
	xlFile.SetCellValue("Sheet1","D6","FDS")
	xlFile.SetCellValue("Sheet1","E6","42")

	err = xlFile.WriteTo("d:/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
