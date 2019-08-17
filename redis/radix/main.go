package main

import (
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/mediocregopher/radix.v2/pool"
	"fmt"
	"sync"
)

func main() {
	commontest()
	//pooltest()
}

func commontest() {
	clent, err := redis.Dial("tcp", "192.168.130.201:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	resq := clent.Cmd("SET", "hello", "world!")
	fmt.Println(resq.Str())
	resq = clent.Cmd("GET", "hello")
	fmt.Println(resq.String())
}
func pooltest() {
	size := 10
	pl, err := pool.New("tcp", "192.168.130.201:6379", size)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	for i := 0; i < size * 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				conn, err := pl.Get()
				if err != nil {
					fmt.Println(err)
					return
				}
				reqp := conn.Cmd("ECHO", "HI")
				fmt.Println(reqp)

				pl.Put(conn)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(pl.Avail())
	pl.Empty()
	fmt.Println(pl.Avail())

}