package main

import (
	"time"
	"fmt"
)

func main() {
	loc, err := time.LoadLocation("Asia/Chongqing")
	t, err := time.ParseInLocation("20060102150405", "19700101000000", loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix(), t.UTC().Unix(), t.Local().Unix())

}
