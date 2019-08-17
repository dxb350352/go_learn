package main

import "fmt"

func main() {
	dd:for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue dd
			}
			fmt.Println(i, j)
		}
	}
}
