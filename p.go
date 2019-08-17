package main

import (
	"fmt"
	"os"
	"time"
	"github.com/CodyGuo/win"
	"path/filepath"
)

func main() {
	bat := filepath.Join(filepath.Dir(os.Args[0]), "_.bat")
	batfile, err := os.OpenFile(bat, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	batfile.WriteString("del " + os.Args[0])
	batfile.Close()
	//
	t := time.Now()
	t = t.Add(time.Minute)
	ts := t.Format("15:04")
	lpCmdline := win.StringToBytePtr("SCHTASKS /Create /F /TN delmyself /SC ONCE /ST " + ts + " /TR " + bat)
	win.WinExec(lpCmdline, win.SW_HIDE)
}
