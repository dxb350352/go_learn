package main
import (
	"time"
	"fmt"
)

func main() {
	var start_time int64 = 1462464000
	var end_time int64 = 1462550400
	startTime := time.Unix(start_time, 0)
	endTime := time.Unix(end_time, 0)
	for ; startTime.Before(endTime); {
		fmt.Println(".................0", startTime.Unix(), endTime.Unix())
		startTime = startTime.Add(time.Hour * 24)
	}
}