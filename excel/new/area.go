package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"strings"
	"github.com/mozillazg/go-pinyin"
	"regexp"
)

//简称
var jiancheng = map[string]string{}

func init() {
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
			code, err := row.Cells[0].String()
			if err != nil {
				continue
			}
			name, err := row.Cells[3].String()
			if err != nil {
				continue
			}
			jiancheng[code] = name
		}
	}
}

var lawreq = regexp.MustCompile(`\(.*\)`)
//http://202.108.98.30/fuzzySearch条件输入%可以查出所有
//把表格内数据copy到excel，直辖市多加一行,go run
func main() {
	xfile2 := xlsx.NewFile()
	sheet2, err := xfile2.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	//第一行
	row2 := sheet2.AddRow()
	row2.AddCell().SetValue("100000")
	row2.AddCell().SetValue("中国")
	row2.AddCell().SetValue("0")
	row2.AddCell().SetValue("中国")
	row2.AddCell().SetValue("0")
	row2.AddCell().SetValue("")
	row2.AddCell().SetValue("")
	row2.AddCell().SetValue("中国")
	row2.AddCell().SetValue("")
	row2.AddCell().SetValue("")
	row2.AddCell().SetValue("china")
	row2.AddCell().SetValue("")

	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/new/china.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			province, _ := row.Cells[0].String()
			if province == "" {
				continue
			}
			city, _ := row.Cells[1].String()
			city = strings.Replace(city, "☆", "", -1)
			if strings.Contains(city, "直辖县级") {
				city = "省直辖县"
			}
			district, _ := row.Cells[2].String()
			areacode, _ := row.Cells[3].String()
			if areacode == "" {
				areacode, _ = sheet.Rows[i+1].Cells[3].String()
				areacode = areacode[:4] + "00"
			}
			phonecode, _ := row.Cells[4].String()
			mailcode, _ := row.Cells[5].String()
			lower := lawreq.FindAllString(province, 1)[0]
			province = lawreq.ReplaceAllString(province, "")
			province = strings.TrimRight(province, "市")
			arealaw := TrimLoopSuffix(areacode, "00")
			level := len(arealaw) / 2
			row2 := sheet2.AddRow()
			row2.AddCell().SetValue(areacode)
			current := district
			if current == "" {
				current = city
			}
			if current == "" {
				current = province
			}
			row2.AddCell().SetValue(current)
			parent := "100000"
			if level > 1 {
				parent = arealaw[:len(arealaw)-2]
				for len(parent) < 6 {
					parent += "0"
				}
			}
			row2.AddCell().SetValue(parent)
			short := jiancheng[areacode]
			if short == "" && len(current) > 2 {
				short = strings.TrimRight(current, "省市区县")
			}
			row2.AddCell().SetValue(short)
			row2.AddCell().SetValue(level)
			if phonecode == "" {
				row2.AddCell().SetValue("")
			} else {
				row2.AddCell().SetValue("0" + phonecode)
			}
			row2.AddCell().SetValue(mailcode)
			all := []string{"中国", province}
			if city != "" {
				all = append(all, city)
			}
			if district != "" {
				all = append(all, district)
			}
			row2.AddCell().SetValue(strings.Join(all, ","))
			row2.AddCell().SetValue("")
			row2.AddCell().SetValue("")
			row2.AddCell().SetValue(strings.Join(pinyin.LazyConvert(short, nil), ""))

			lower = strings.Replace(lower, "(", "", -1)
			lower = strings.Replace(lower, ")", "", -1)
			row2.AddCell().SetValue(lower)
		}
	}

	err = xfile2.Save("E:/GOPATH/src/testgo/excel/new/area1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

//循环去掉后缀知道没有为止
func TrimLoopSuffix(s, suffix string) string {
	for {
		if strings.HasSuffix(s, suffix) {
			s = strings.TrimSuffix(s, suffix)
		} else {
			return s
		}
	}
	return ""
}
