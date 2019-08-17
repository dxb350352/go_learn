package main
import (
	"time"
	"fmt"
	"log"
)

func main() {
	//http://studygolang.com/articles/2593
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())


	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	fmt.Println("start...............")
	pp()
	fmt.Println("after panic.........")
	time.Sleep(time.Hour);
}

func pp() {
	log.Panic("pp panic")
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 2
	defer func() {
		t = t + 2
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}