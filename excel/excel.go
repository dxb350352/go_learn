package main
import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"strings"
)

func main() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/ss.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var areas string = "{"
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			name,_:=row.Cells[7].String()
			lnt,_:=row.Cells[8].String()
			lat,_:=row.Cells[9].String()
			areas += fmt.Sprint(`"`, strings.TrimLeft(name,"中国,"), `":[`, lnt, `,`, lat, `],`)
		}
	}
	temp := []byte(areas)
	temp[len(temp) - 1] = ([]byte(`}`))[0]

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