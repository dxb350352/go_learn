package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"strings"
)

//简称
var jiancheng = map[string]string{}
var phone8 = map[string]string{}

func init() {
	initPhone8()
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/ss.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 4 {
				continue
			}
			fullname, err := row.Cells[1].String()
			if err != nil {
				continue
			}
			name, err := row.Cells[3].String()
			if err != nil {
				continue
			}
			jiancheng[fullname] = name
		}
	}
}

func initPhone8() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/new/phone8.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 4 {
				continue
			}
			code, err := row.Cells[1].String()
			if err != nil {
				continue
			}
			num, err := row.Cells[3].String()
			if err != nil {
				continue
			}
			if num == "8" {
				phone8[code] = num
			}
		}
	}
}

func main() {
	xfile2 := xlsx.NewFile()
	sheet2, err := xfile2.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	//
	m := map[string]bool{}
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/new/area.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 8 {
				continue
			}
			phone, err := row.Cells[5].String()
			if err != nil {
				continue
			}
			code, _ := row.Cells[0].String()
			//台湾
			if strings.HasPrefix(code, "71") {
				phone = "00886" + phone
			}
			if len(phone) > 0 && !m[phone] {
				name, _ := row.Cells[7].String()
				name = strings.TrimPrefix(name, "中国,")
				if !strings.Contains(name, ",") {
					name = name + "," + name
				}
				namearr := strings.Split(name, ",")
				//简称
				for i, v := range namearr {
					if jiancheng[v] == "" {
						if len(v) > 2 {
							namearr[i] = strings.TrimRight(v, "省市区县")
						}
					} else {
						namearr[i] = jiancheng[v]
					}
				}
				name = strings.Join(namearr, ",")
				row2 := sheet2.AddRow()
				row2.AddCell().SetValue(phone)
				row2.AddCell().SetValue(code)
				row2.AddCell().SetValue(name)
				row2.AddCell().SetValue(phone8[code])
				m[phone] = true
			}
		}
	}

	err = xfile2.Save("E:/GOPATH/src/testgo/excel/new/phone1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
