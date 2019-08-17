package main

import (
	"sync"
	"fmt"
	"time"
)

var Condition int
var Lock = new(sync.Mutex)
var Cond = sync.NewCond(Lock)

func main() {
	for {
		Lock.Lock()
		Condition++
		for Condition > 10 {
			fmt.Println(".........")
			Cond.Wait()
		}
		go test()
		Lock.Unlock()
	}
	time.Sleep(time.Second * 10)
}

func test() {
	fmt.Println("test.......start", Condition)
	time.Sleep(time.Second)
	Condition--
	fmt.Println("test.......end", Condition)
	Cond.Signal()
}