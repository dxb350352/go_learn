package main

import (
	"io"
	"crypto/md5"
	"fmt"
	"strings"
	"crypto/sha1"
)

func main() {
	md5go()
}

func md5go() {
	password := "unzip_pwd101-001"
	h := md5.New()
	io.WriteString(h, password)

	md := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(md)
}

func sha1go() {
	serverSHA1Param := fmt.Sprint("user:", 1501567658, ":THISISAPPSALT")
	//h := sha1.New()
	h := sha1.New()
	h.Write([]byte(serverSHA1Param))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	sha1str := fmt.Sprintf("%02x", bs)
	fmt.Println(sha1str)
	var bsstr [] string
	for _, v := range bs {
		tmp := fmt.Sprintf("%02x", v & 0xFF)
		if len(tmp) < 2 {
			bsstr = append(bsstr, "0")
		}
		bsstr = append(bsstr, tmp)
	}
	fmt.Println(strings.Join(bsstr, ""))
}
