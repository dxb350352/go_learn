package main

import (
	"path/filepath"
	"github.com/extrame/xls"
	"log"
	"fmt"
	"regexp"
)
var blank = regexp.MustCompile(`\s`)
func main() {

	xlFile, err := xls.Open(filepath.Join( "E:/GOPATH/src/testgo/excel/User.xls"), "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < xlFile.NumSheets(); i++ {
		sheet := xlFile.GetSheet(i)
		if sheet == nil {
			continue
		}
		maxrow := int(sheet.MaxRow)
		for j := 0; j <= maxrow; j++ {
			row := sheet.Row(j)
			if row == nil {
				continue
			}
			fmt.Println(row.Col(4))
			fmt.Println(getExcelValue(row.Col(4)))
		}

	}
}

func getExcelValue(str string)  string{
	r:=[]rune(str)
	for i:=0;i<len(r);i++{
		if r[i]==0{
			r=append(r[:i],r[i+1:]...)
			i--
		}
	}
	return string(r)
}