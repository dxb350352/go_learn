package main

import (
	"fmt"
	"time"
	"github.com/jakecoffman/cron"
)

type name struct {
	Id int
}

func main() {
	c := cron.New()
	c.Start()
	var n name
	c.AddFunc("0/5 * * * * ?", func() {
		n.Id += 1
		fmt.Println(n.Id)
	}, "ttt")
	time.Sleep(time.Hour)
}
