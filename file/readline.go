package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	arr, err := ReadFileCount("f:/new1.txt", 3)
	arr, err = ReadFileCount("f:/new1.txt", 13)
	fmt.Println(err)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	//file, err := os.OpenFile("D:/Desktop/文档/No.5/data/001-002.txt", os.O_RDONLY, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//buf := bufio.NewReader(file)
	//for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
	//	line, isPrefix, err = buf.ReadLine()
	//	fmt.Println(string(line), isPrefix, err)
	//}
	////byt,err:=ioutil.ReadFile("D:/Desktop/文档/No.5/data/001-001.txt")
	////if err != nil {
	////	log.Fatal(err)
	////}
	////fmt.Println(string(byt))
}

//读取文件一定行数
func ReadFileCount(path string, count int64) ([]string, error) {
	var result []string
	if count <= 0 {
		return result, nil
	}
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return result, err
	}
	defer file.Close()
	iscan := bufio.NewScanner(file)
	var n int64
	for iscan.Scan() {
		line := iscan.Text()
		result = append(result, line)
		n++
		if n >= count {
			break
		}
	}
	return result, nil
}
