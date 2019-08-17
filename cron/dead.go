package main

import (
	"github.com/jakecoffman/cron"
	"time"
	"fmt"
	"runtime"
	"strings"
)

func main() {
	c := cron.New()
	c.Start()
	c.AddFunc("0/5 * * * * ?", func() {
		fmt.Println(GoroutineID())

	}, "ttt")
	time.Sleep(time.Hour)
}

func GoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	return idField
}
