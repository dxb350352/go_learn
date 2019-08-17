package main

import (
	"time"
	"fmt"
	"sync"
)
var waitgrou sync.WaitGroup
func main() {
	after()
}

func timer() {
	timer := time.NewTimer(1 * time.Second)

	<-timer.C
	fmt.Println("Timer has expired.")

	timer.Stop()
}

func ticker() {
	ticker := time.NewTicker(time.Second * 2)
	for i := range ticker.C {
		fmt.Println(i)
		ticker.Stop()
		break
	}
}

func ticker2() {
	ticker := time.NewTicker(time.Second * 2)
	<-ticker.C
	fmt.Println("over")
	ticker.Stop()

}

func after() {
	fmt.Println(time.Now(),111)
	waitgrou.Add(1)
	time.AfterFunc(time.Second,func(){
		fmt.Println(time.Now(),222)
		waitgrou.Done()
	})
	waitgrou.Wait()
}

func after2() {
	fmt.Println(".............")
	<-time.After(time.Second)
	fmt.Println(1234)
}