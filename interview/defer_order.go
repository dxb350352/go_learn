package main

import "fmt"

//defer 是后进先出。
//panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=", err)
		}
	}()

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
