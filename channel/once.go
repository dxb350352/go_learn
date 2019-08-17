package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	twoprint()
	time.Sleep(time.Second*3)
}

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}
func doprint() {
	once.Do(setup)
	fmt.Println(a)
}
func twoprint() {
	go doprint()
	go doprint()
}
