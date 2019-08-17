package main

import (
	"os"
	"fmt"
	"github.com/hydra13142/chardet"
)

func main() {
	file, err := os.Open("D:/Desktop/文档/No.5/话单/用户详单15641765017-5.xls")
	if err != nil {
		fmt.Println(err)
		return
	}
	byt := make([]byte, 10240)
	i, err := file.Read(byt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	arr := chardet.Possible(byt)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println(chardet.Mostlike(byt))
}
