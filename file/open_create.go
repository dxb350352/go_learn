package main

import (
	"os"
	"fmt"
	"path/filepath"
)
//结论是没有上层目录，都不能创建嘛
func main() {
	path1 := "d:/test1/1.txt"
	path2 := "d:/test2/2.txt"
	path3 := "d:/test3/3.txt"
	_, err := os.OpenFile(path1, os.O_CREATE | os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = os.Create(path2)
	if err != nil {
		fmt.Println(err.Error())
	}
	path3, fname := filepath.Split(path3)
	err = os.MkdirAll(path3, os.ModeDir)
	if err != nil {
		fmt.Println(err.Error())
	}
	path3 = filepath.Join(path3, fname)
	_, err = os.OpenFile(path3, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}