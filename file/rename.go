package main

import (
	"os"
	"fmt"
)

func main() {
	src:="d:/FeiQ.exe"
	des:="d:/1111/123"
	err:=os.Rename(des,src)
	fmt.Println(err,"............err")
}
