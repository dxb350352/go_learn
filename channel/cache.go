package main
import "fmt"

func main() {
	chs := make(chan int, 10240)
	for i := 0; i < 10240; i++ {
		fmt.Println("put", i)
		chs <- i
	}
	//	for v := range chs {
	//		fmt.Println(v)
	//	}
	for i := 0; i < 10240; i++ {
		fmt.Println("get", <-chs)
	}
}
