package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func main() {
	write3()
}

func write3() {
	xfile1, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/statis_model.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	sheet1 := xfile1.Sheets[0]
	xfile2 := xlsx.NewFile()
	sheet2, err := xfile2.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	row:=sheet2.AddRow()
	cell:=row.AddCell()
	cell.SetValue(123)
	cell.SetStyle(sheet1.Cell(2,1).GetStyle())
	fmt.Println(sheet1.Cell(2,1).Value)
	err = xfile2.Save("d:/test.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func write2() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/statis_model.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	var file *xlsx.File

	file = xlsx.NewFile()
	_, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	file.Sheets[0] = xlFile.Sheets[0]

	file.Sheets[0].Cell(1, 0).SetValue(file.Sheets[0].Cell(1, 0).Value + "fdsfsdfs")
	fmt.Println(file.Sheets[0].Cell(1, 0).Value)

	file.Sheets[0].Cell(1, 21).SetValue(file.Sheets[0].Cell(1, 21).Value + "fdsfsdfs")
	fmt.Println(file.Sheets[0].Cell(1, 21).Value)

	err = file.Save("d:/test.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func write1() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	err = file.Save("d:/test.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func write() {
	xlFile, err := xlsx.OpenFile("E:/GOPATH/src/testgo/excel/statis_model.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	xlFile.Sheets[0].Cell(1, 0).SetValue(xlFile.Sheets[0].Cell(1, 0).Value + "fdsfsdfs")
	fmt.Println(xlFile.Sheets[0].Cell(1, 0).Value)

	xlFile.Sheets[0].Cell(1, 21).SetValue(xlFile.Sheets[0].Cell(1, 21).Value + "fdsfsdfs")
	fmt.Println(xlFile.Sheets[0].Cell(1, 21).Value)

	err = xlFile.Save("d:/test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}
