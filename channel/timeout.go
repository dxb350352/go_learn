package main
import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()
	select {
	case a := <-ch:
		fmt.Println(a)
	case b := <-timeout:
		fmt.Println(b)
	}
	fmt.Println("OK")
}
