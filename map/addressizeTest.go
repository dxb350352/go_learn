package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	models "testgo/map/copy"
	"time"
	"strings"
)

func main() {
	initAreaMap()
	s1 := time.Now().UnixNano()
	//str := "北京四川省四川通州区汤立路201号奥区3-1-29-01号"
	//str := "广东省,东山区新阳花园白领商贸,东山"
	//str := "其他,长沙五一路368号,湖南 长沙"
	//str := "北京,嘉友国际大厦北区"
	//str := "四川省四川省成都市青羊区成都"
	//str := "上海市陕西北路2342-32-34"
	//str := "广西壮族自治区玉林市，新民街29号"
	//str := "上海市，上海市长宁区江苏路121-123"
	//str := "新会大泽新兴区鑫源"
	//str := "上海市，上海市长宁区江苏街121-123"
	//str := "fdsfa"
	//str := "中国山西临汾go"
	//str := "湖北省武汉市"
	//str := ""
	//str := ","
	//str := "四川省成都市青羊區天府((((((((,,,,,,,)))))))))大道１２３號軟件園A區"
	//str := "澳门黑沙环广福安花园"
	//str := "呼和浩特市新城区中山里街美新花园11号楼9单元501号"
	//str := "内蒙古满洲里市道北三道街富豪小区3号楼9单元501号"
	str := " 142730"
	fmt.Println(str)
	for i := 0; i < 1; i++ {
		//fmt.Println("...............")
		fmt.Println(models.GetAddressValidated(str))
		//models.GetAddressValidated(str)
	}
	fmt.Println(time.Now().UnixNano() - s1)
	//file, err := os.Open("E:/GOPATH/src/testgo/map/copy/test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//arr, err := utils.ReadByFileCount(file, 300)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for i, v := range arr {
	//	av := strings.Split(v, "|")
	//	fmt.Println(i, "--", models.GetAddressValidated(av[2]), "--", av[2])
	//	fmt.Println(i, "--", models.GetAddressValidated(av[6]), "--", av[6])
	//}

}

//初始化地址解析分词器
func initAreaMap() {
	models.AreaMap = make(map[rune]interface{})
	models.AreaMapShort = make(map[rune]interface{})
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/github.com/sas/dataimport/conf/area.xlsx")
	if err != nil {
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) > 7 {
				models.Put(models.AreaMap, []rune(strings.TrimSpace(row.Cells[1].Value)), strings.TrimSpace(row.Cells[7].Value))
				models.Put(models.AreaMapShort, []rune(strings.TrimSpace(row.Cells[3].Value)), strings.TrimSpace(row.Cells[7].Value))
			}
		}
	}
}
