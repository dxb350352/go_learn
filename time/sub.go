package main

import (
	"time"
	"fmt"
)

func main() {
	t1,_:=time.Parse("2006-01-02","2017-11-29")
	t2,_:=time.Parse("2006-01-02","2017-11-28")
	fmt.Println(int64(t1.Sub(t2).Hours()/24))
}
