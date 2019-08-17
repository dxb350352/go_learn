package main

import "fmt"

func main() {
	lines := []int{1, 2, 3, 4, 5, 6, 7}
	var count int = 3
	for i, v := range lines {
		if i < count {
			continue
		}
		fmt.Println(lines[i - count + 1 :i])
		fmt.Println(v)
	}
}
