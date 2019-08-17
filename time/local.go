package main

import (
	"fmt"
	"time"
)

func main() {
	//st := "2016-11-12 18:00"
	//fmt.Println(st[5 : len(st)-1-5])
	//time3, _ := time.Parse("2006-01-02 15:04", st)
	//loc, _ := time.LoadLocation("Local")
	//fmt.Println(loc)
	//time2, _ := time.ParseInLocation("2006-01-02 15:04", st, loc)
	//fmt.Println(time3)
	//fmt.Println(time2)
	testlocal()
}

func testlocal()  {
	st:="1970-02"
	t,err:=time.Parse("2006-01",st)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(t.Unix())
	}
}