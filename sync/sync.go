package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("...........start")
	var wg sync.WaitGroup
	wg.Add(1)
	go ss(&wg)
	wg.Wait()
	fmt.Println("...........end")
}

func ss(wg *sync.WaitGroup) {
	fmt.Println("...........mid")
	wg.Done()
}