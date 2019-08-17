package main

import (
	"fmt"
)

var ch chan int = make(chan int,100)

func foo(id int) { //id: 这个routine的标号
	fmt.Println(id,"....")
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 50; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 50; i++ {
		fmt.Println(<- ch)
	}
}
