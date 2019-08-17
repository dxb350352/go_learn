package main

import "fmt"
import "time"

func send(ch1 chan int, ch2 chan int) {
	i := 0
	for {
		i++
		select {
		case ch1 <- i:
			fmt.Printf("send ch1 %d\n", i)
		case ch2 <- i:
			fmt.Printf("send ch2 %d\n", i)
		default:
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func recv(ch chan int, gap time.Duration, name string) {
	for v := range ch {
		fmt.Printf("receive %s %d\n", name, v)
		time.Sleep(gap)
	}
}

//非阻塞读写
//生产者同时向两个通道写数据，写不进去就丢弃。
func main() {
	// 无缓冲通道
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	// 两个消费者的休眠时间不一样，名称不一样
	go recv(ch1, time.Second, "ch1")
	go recv(ch2, 2*time.Second, "ch2")
	send(ch1, ch2)
}
