package main

import (
	"sync"
	"time"
)

var m = make(map[string]string)
var l = sync.Mutex{}

func main() {
	go func() {
		for {
			//l.Lock()
			m["x"] = "foo"
			//l.Unlock()
		}
	}()
	go func() {
		for {
			//l.Lock()
			m["x"] = "foo"
			//l.Unlock()
		}
	}()

	time.Sleep(3 * time.Second)
}