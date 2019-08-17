package main

import (
	"github.com/CodyGuo/win"
	"fmt"
)

func main() {
	lpCmdline := win.StringToBytePtr("setx aa11 abcd")
	ret := win.WinExec(lpCmdline, win.SW_HIDE)
	fmt.Println(ret)

}
