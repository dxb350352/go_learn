package main

import "fmt"

func main() {
	var msgs = make(chan string)
	go func(msg string) {
		msgs <- msg
	}("hehe")

	fmt.Println(<-msgs)
}
