package main

import (
	"testgo/excel/source_code/optimize2/xlsx"
	"fmt"
	"unsafe"
	"syscall"
	"time"
	"os"
	"archive/zip"
)

//这个优化针对sheet会被修改
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
		writeToFile(file, zipWriter, i)
	}
	writeToFile(file, zipWriter, -2)
	flag = false
}

func writeToFile(file *xlsx.File, zipWriter *zip.Writer, i int) {
	parts, err := file.MarshallParts(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range parts {
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
