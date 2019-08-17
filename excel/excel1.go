package main

import (
	"fmt"
	"github.com/extrame/xls"
	"strings"
	"time"
)
var DWDM = map[string]string{}
func main() {
	start := time.Now().UnixNano() / 1e6
	xlFile, err := xls.Open("E:/GOPATH/src/testgo/excel/DWDM.xls", "utf-8")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1; i++ {
		sheet := xlFile.GetSheet(i)
		if sheet == nil {
			break
		}
		//8455 江西省新余市公安消防支队后勤处财务科
		for j := 8455; j < 8456; j++ {
			row := sheet.Row(j)
			if row == nil {
				break
			}
			if row.LastCol() < 1 {
				continue
			}
			dm := strings.TrimSpace(row.Col(0))
			if dm == "" {
				continue
			}

			mc := strings.TrimSpace(row.Col(1))
			if dm == "" {
				continue
			}
			fmt.Println(dm, "-", mc)
			DWDM[dm] = mc
		}
	}
	fmt.Println(time.Now().UnixNano()/1e6 - start)
}