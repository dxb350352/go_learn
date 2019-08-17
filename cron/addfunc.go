package main

import (
	"github.com/jakecoffman/cron"
	"fmt"
	"time"
)

func main() {
	c := cron.New()
	c.Start()
	c.AddFunc("0/2 * * * *", func() {
		fmt.Println(time.Now(), "..........1")
	}, "test1")
	c.AddFunc("5/10 * * * * *", backup, "alertarea" )
	//time.Sleep(time.Second * 5)
	//c.AddFunc("0/61 * * * *", backup, "test2")
	time.Sleep(time.Hour)
}

func backup() {
	fmt.Println(time.Now(), "..........2")
}