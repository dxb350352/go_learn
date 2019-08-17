package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go print(i, &wg)
	}
	wg.Wait()
}

func print(i interface{}, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}


