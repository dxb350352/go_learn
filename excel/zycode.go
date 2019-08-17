package main

import (
	"fmt"
	"os"
	"github.com/tealeg/xlsx"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/zhiye.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var areas string = "["
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			v := row.Cells[0].Value
			areas += "{"
			areas += fmt.Sprint(`"id":"`, v, `",`)
			areas += fmt.Sprint(`"name":"`, row.Cells[1].Value, `",`)
			areas += fmt.Sprint(`"open":`, len(v) < 2, `,`)
			areas += fmt.Sprint(`"pId":"`, v[:len(v) - 1], `"`)
			areas += "},"
		}
	}
	temp := []byte(areas)
	temp[len(temp) - 1] = ([]byte(`]`))[0]

	fmt.Println(string(temp))

	file, err := os.Create(`E:/GOPATH_GO/src/testgo/excel/zhiyecode.json`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}