package main

import (
	"os"
	"crypto/md5"
	"io"
	"fmt"
)

func main() {
	fmt.Println(md5file("C:/Users/Administrator/Desktop/dataimport.log.2017.09.15"))
}

func md5file(fpath string) string {
	file, err := os.Open(fpath)
	if err != nil {
		return ""
	}
	defer file.Close()
	md5mash := md5.New()
	_, err = io.Copy(md5mash, file)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5mash.Sum(nil))
}