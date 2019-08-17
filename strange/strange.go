package main
import "fmt"

func main() {
	//这TM是“或”运算啊
	test(1 | 2 | 4)
}

func test(a int) {
	fmt.Println(a)
}
