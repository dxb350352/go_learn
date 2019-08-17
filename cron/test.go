package main

import (
	"github.com/jakecoffman/cron"
	"fmt"
	"time"
	"github.com/sas/utils"
)

func main() {
	c := cron.New()
	c.Start()
	LogsSyncTime := 1
	SychronizeCron := fmt.Sprintf("0 0/%d * * * *", LogsSyncTime)
	if utils.CronValidate(SychronizeCron) {
		c.AddFunc(SychronizeCron, backup1, "area")
	}
	time.Sleep(time.Hour)
}

func backup1() {
	fmt.Println(time.Now(), "..........2")
}