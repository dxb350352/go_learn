package main

import (
	"runtime"
	"sync"
	"fmt"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	main111()
	main222()
}

//闭包
func main111() {
	var f = Adder()
	fmt.Println(f(1), "-")
	fmt.Println(f(20), "-")
	fmt.Println(f(300), "-")

}
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

//闭包
func main222() {
	//判断字符串 以bmp结尾
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))
	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
