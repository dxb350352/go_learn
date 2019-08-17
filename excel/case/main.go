package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	score()
}

func basecode() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/case/自定义类别编码20180117-结果.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			name, _ := row.Cells[1].String()
			code, _ := row.Cells[2].String()
			fmt.Println(name, code)
		}
	}
}

func score() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/case/评分模型合集.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//mode1
	for i, row := range xlFile.Sheet["mode1"].Rows {
		if i == 0 {
			continue
		}
		name, _ := row.Cells[0].String()
		code, _ := row.Cells[1].String()
		fmt.Println(name, code)
	}
	//mode2
	for i, row := range xlFile.Sheet["mode2"].Rows {
		if i == 0 {
			continue
		}
		name, _ := row.Cells[0].String()
		code, _ := row.Cells[2].String()
		fmt.Println(name, code)
	}
	//mode3
	for i, row := range xlFile.Sheet["mode3"].Rows {
		if i == 0 {
			continue
		}
		name, _ := row.Cells[0].String()
		code, _ := row.Cells[2].String()
		fmt.Println(name, code)
	}
	//mode4
	for i, row := range xlFile.Sheet["mode4"].Rows {
		if i == 0 {
			continue
		}
		name, _ := row.Cells[0].String()
		code, _ := row.Cells[2].String()
		fmt.Println(name, code)
	}
}
