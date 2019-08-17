package main
import "fmt"

func main() {
	if false && test1() {
		fmt.Println("1")
	}
}

func test1() bool {
	fmt.Println("2")
	return true
}
