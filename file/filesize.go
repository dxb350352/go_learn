package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	f,err:=os.Open("E:/迅雷下载/arangodb-linux-amd64")
	if err != nil {
		fmt.Println(err)
		return
	}
	finfo,err:=f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(finfo.Size())

	fd,err := ioutil.ReadAll(f)
	fmt.Println(len(fd))
}
