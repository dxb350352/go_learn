package main

import (
	"log"
	"github.com/aswjh/excel"
	"fmt"
)

func main() {
	xl, err := excel.Open("E:/GOPATH/src/testgo/excel/bill/话单.xls", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer xl.Quit()
	for i := 0; i < xl.CountSheets(); i++ {
		sheet := xl.Sheets()[i]
		start := "A"
		end := string('A' + getColCount(sheet) - 1)
		fmt.Println(end)
		for r := 1; r < 65536; r++ {
			ret, err := sheet.GetRange(fmt.Sprintf("%s%d:%s%d", start, r, end, r))
			if err != nil {
				break
			}
			row := ret.([][]interface{})
			if len(row) == 0 || len(row[0]) == 0 || len(row[0][0].(string)) == 0 {
				break
			}
			var cols []string
			for _, col := range row[0] {
				cols = append(cols, col.(string))
			}
			fmt.Println(cols)
		}
	}
}

func getColCount(sheet excel.Sheet) int {
	for i := 1; i < 100; i++ {
		ret, err := sheet.Cells(1, i)
		if err != nil || ret == "" {
			return i - 1
		}
	}
	return 0
}
