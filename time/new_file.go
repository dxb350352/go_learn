package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {
	var t time.Time
	fmt.Println(t.Unix())
	t=time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("20060102150405"))


	typet:=reflect.TypeOf(t)
	fmt.Println(typet.Name())

	fmt.Println(time.Unix(1456193459,0))
	fmt.Println(time.Now().Day(),".......")


	newday, _ := time.Parse("2006-01-02",time.Now().Format("2006-01-02"))
	fmt.Println(newday)
	fmt.Println(newday.Add(-time.Second))
	fmt.Println(newday)

}
