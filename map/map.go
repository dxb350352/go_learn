package main
import (
	"fmt"
	"time"
)

var m map[int]int = map[int]int{
	0:0,
	1:1,
	2:2,
	3:3,
	4:4,
	5:5,
	6:6,
	7:7,
	8:8,
	9:9,
}
func main() {
	m1 := map[int]int{
		1:1,
		2:2,
	}
	m2 := map[int]int{
		3:3,
		4:4,
	}
	fmt.Println(m1, m2)
	go add()
	go remove()
	time.Sleep(time.Second * 5)
	fmt.Println(m)
}

func remove() {
	for k, _ := range m {
		if k >= 10 {
			delete(m, k)
			fmt.Println("remove", k, m)
		}
	}
}

func add() {
	for k, v := range m {
		m[k * 10] = v * 10
		fmt.Println("add", k, m)
	}
}