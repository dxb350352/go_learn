package main

import (
	"testgo/excel/source_code/optimize1/xlsx"
	"fmt"
	"unsafe"
	"syscall"
	"time"
	"os"
	"archive/zip"
	"strings"
)

//这个优化针对sheet不会被修改
func main() {
	flag := true
	go func() {
		for flag {
			printMem()
			time.Sleep(time.Second)
		}
	}()

	writer, err := os.Create("E:/big.xlsx")
	if err != nil {
		return
	}
	defer writer.Close()
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	refTable := xlsx.NewSharedStringRefTable()

	file := xlsx.NewFile()
	for i := 0; i < 2000; i++ {
		sheet, err := file.AddSheet(fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
			continue
		}
		for ii := 0; ii < 100; ii++ {
			row := sheet.AddRow()
			for iii := 0; iii < 100; iii++ {
				row.AddCell().SetValue("1234567890")
			}
		}
		writeToFile(file, zipWriter, refTable, i)
	}
	writeToFile(file, zipWriter, refTable, -1)
	flag = false
}

func writeToFile(file *xlsx.File, zipWriter *zip.Writer, refTable *xlsx.RefTable, i int) {
	parts, err := file.MarshallParts(refTable)
	if err != nil {
		return
	}
	for k, v := range parts {
		sheetNameIndex := i + 1
		if sheetNameIndex < 0 && !strings.HasPrefix(k, "xl/worksheets/") || sheetNameIndex > 0 && k == fmt.Sprintf("xl/worksheets/sheet%d.xml", sheetNameIndex) {
			fwriter, err := zipWriter.Create(k)
			if err != nil {
				return
			}
			_, err = fwriter.Write([]byte(v))
			if err != nil {
				return
			}
			fmt.Println(k)
		}
	}
	if i >= 0 {
		sheet := &xlsx.Sheet{Name: fmt.Sprint(i), File: file}
		file.Sheets[i] = sheet
		file.Sheet[sheet.Name] = sheet
	}
}

var kernel = syscall.NewLazyDLL("Kernel32.dll")

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func printMem() {
	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return
	}
	fmt.Println("total=:", memInfo.ullTotalPhys/1024/1024, "free=:", memInfo.ullAvailPhys/1024/1024)
}
