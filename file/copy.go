package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	src := "e:/test1.txt"
	des := "e:/test2.txt"
	copy1(src, des)
	//copy2(src, des)
}
func copy1(srcstr, desstr string) {
	src, err := os.Open(srcstr)
	if err != nil {
		fmt.Println(1, err)
	}
	defer src.Close()
	//dst, err := os.OpenFile("e:/test2.txt", os.O_CREATE | os.O_TRUNC, 0666)
	dst, err := os.Create(desstr)
	if err != nil {
		fmt.Println(2, err)
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println(3, err)
	}
	//err = dst.Sync()
	//if err != nil {
	//	fmt.Println(4, err)
	//}
}

func copy2(src, des string) {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return
	}
	if _, err := os.Stat(des); !os.IsNotExist(err) {
		os.Remove(des)
	}
	err := os.Link(src, des)
	if err != nil {
		fmt.Println(5, err)
	}
}
