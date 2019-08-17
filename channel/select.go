package main

import "fmt"

import (
	"time"
	"os"
)

func launch() {
	fmt.Println("nuclear launch detected")
}

func commencingCountDown(canLunch chan int) {
	c := time.Tick(time.Second)
	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-c
	}
	canLunch <- -1
}

func isAbort(abort chan int) {
	byt := make([]byte, 1)
	os.Stdin.Read(byt)
	abort <- -1
}

func main() {
	fmt.Println("Commencing coutdown")

	abort := make(chan int)
	canLunch := make(chan int)
	go isAbort(abort)
	go commencingCountDown(canLunch)
	select {
	case a := <-canLunch:
		fmt.Println(a)
	case a := <-abort:
		fmt.Println("Launch aborted!", a)
		return
	}
	launch()
}
