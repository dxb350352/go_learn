package main

import "fmt"

func main() {
	lines := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var count int = 3
	total := len(lines) / count
	for i := 0; i < total; i++ {
		fmt.Println(lines[i * count: (i + 1) * count])
	}
	fmt.Println(lines[total * count:])
}
