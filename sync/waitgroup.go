package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("...........start")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sss(&wg)
	}()
	wg.Wait()
	fmt.Println("...........end")
}

func sss(wg *sync.WaitGroup) {
	fmt.Println("...........mid")
	wg.Done()

}