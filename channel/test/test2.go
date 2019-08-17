package main

import (
	"fmt"
	"time"
)

var complete = make(chan int)

func loog() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second / 100)
	}
	complete <- 1000
}
func main() {
	go loog()
	fmt.Println(<-complete)
}
