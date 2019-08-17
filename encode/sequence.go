package main

import "fmt"

func main() {
	str:="1234567890中文"
	for i:=0;i<len(str);i++{
		fmt.Printf("0x%x\n",str[i])
	}
}
