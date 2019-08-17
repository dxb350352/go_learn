package main

import (
	"fmt"
	"time"
	"github.com/jakecoffman/cron"
)

func main() {
	c := cron.New()
	c.Start()
	c.AddFunc("0/3 * * * *", func() {
		fmt.Println(time.Now(), "..........start")
		time.Sleep(time.Second * 10)
		fmt.Println(time.Now(), "..........over")
	}, "test1")
	time.Sleep(time.Hour)
}
