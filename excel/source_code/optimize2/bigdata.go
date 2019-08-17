package main

import (
	"fmt"
	"unsafe"
	"syscall"
	"time"
	"testgo/excel/source_code/optimize2/xlsx"
	"os"
)

func main() {
	flag := true
	go func() {
		for flag {
			printMem()
			time.Sleep(time.Second)
		}
	}()

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
	}
	fmt.Println("saving.............")
	fmt.Println(file.Save("E:/big1.xlsx"))
	flag = false
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
	persent := memInfo.ullAvailPhys * 100 / memInfo.ullTotalPhys
	if persent <= 5 {
		fmt.Println("内存占用", 100-persent, "%了")
		os.Exit(-1)
	}
}
