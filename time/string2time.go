package main
import (
	"fmt"
	"time"
)

func main() {
	st := "2006-01-02 08:00"
	time3, _ := time.Parse("2006-01-02 15:04", st)
	loc, _ := time.LoadLocation("Local")
	time2, _ := time.ParseInLocation("2006-01-02 15:04", st, loc)
	fmt.Println(time3)
	fmt.Println(time2)
	test()
}

func test() {
	st := "2006-01-02 15:04:05"
	ts := "2016-04-04 02:02:16"
	t, _ := time.Parse(st, ts)
	fmt.Println(t, ".............")
	t = t.AddDate(0, 0, 1)
	fmt.Println(t, ".............")
}
