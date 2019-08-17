package main

import (
	"time"
	"fmt"
	"sync"
)

var L sync.Mutex

func main() {
	go tt2()
	go tt1()
	go tt2()
	time.Sleep(time.Hour)
}

func tt1() {
	L.Lock()
	defer L.Unlock()
	fmt.Println("t1")
	time.Sleep(time.Minute)
}
func tt2() {
	L.Lock()
	defer L.Unlock()
	fmt.Println("t2")
	time.Sleep(time.Second)
}
