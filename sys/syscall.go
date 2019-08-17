package main

import (
	"fmt"
	"syscall"
	"unsafe"
	"strconv"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll");
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot");
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0));
	if int(pHandle) == -1 {
		return;
	}
	Process32Next := kernel32.NewProc("Process32Next");
	for {
		var proc PROCESSENTRY32;
		proc.dwSize = ulong(unsafe.Sizeof(proc));
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
			fmt.Println("------------------------------------------------")
			fmt.Println("ProcessName : " + string(proc.szExeFile[0:50]));
			fmt.Println("ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)));
			fmt.Println("dwSize : " + strconv.Itoa(int(proc.dwSize)));
			fmt.Println("cntUsage : " + strconv.Itoa(int(proc.cntUsage)));
			fmt.Println("th32ProcessID : " + strconv.Itoa(int(proc.th32ProcessID)));
			fmt.Println("th32DefaultHeapID : " + strconv.Itoa(int(proc.th32DefaultHeapID)));
			fmt.Println("th32ModuleID : " + strconv.Itoa(int(proc.th32ModuleID)));
			fmt.Println("cntThreads : " + strconv.Itoa(int(proc.cntThreads)));
			fmt.Println("th32ParentProcessID : " + strconv.Itoa(int(proc.th32ParentProcessID)));
			fmt.Println("pcPriClassBase : " + strconv.Itoa(int(proc.pcPriClassBase)));
			fmt.Println("dwFlags : " + strconv.Itoa(int(proc.dwFlags)));
			fmt.Println("------------------------------------------------")
		}else {
			break;
		}
	}
	CloseHandle := kernel32.NewProc("CloseHandle");
	_, _, _ = CloseHandle.Call(pHandle);
}
