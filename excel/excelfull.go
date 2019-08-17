package main
import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/ss.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var areas string = "["
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			areas += "{"

			areas += fmt.Sprint(`"id":"`, row.Cells[0], `",`)
			areas += fmt.Sprint(`"name":"`, row.Cells[1], `",`)
			areas += fmt.Sprint(`"parent_id":"`, row.Cells[2], `",`)
			areas += fmt.Sprint(`"short_name":"`, row.Cells[3], `",`)
			areas += fmt.Sprint(`"level":"`, row.Cells[4], `",`)
			areas += fmt.Sprint(`"city_code":"`, row.Cells[5], `",`)
			areas += fmt.Sprint(`"zip_code":"`, row.Cells[6], `",`)
			areas += fmt.Sprint(`"merger_name":"`, row.Cells[7], `",`)
			areas += fmt.Sprint(`"longitude":"`, row.Cells[8], `",`)
			areas += fmt.Sprint(`"latitude":"`, row.Cells[9], `",`)
			areas += fmt.Sprint(`"pinyin":"`, row.Cells[10], `"`)

			areas += "},"
		}
	}
	temp := []byte(areas)
	temp[len(temp) - 1] = ([]byte(`]`))[0]

	fmt.Println(string(temp))

	file, err := os.Create(`E:/GOPATH_GO/src/testgo/excel/ss.json`)
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