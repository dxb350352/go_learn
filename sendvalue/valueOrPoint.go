package main
import (
	"fmt"
)
/**
函数参数传值, 闭包传引用!
slice 含 values/count/capacity 等信息, 是按值传递
按值传递的 slice 只能修改values指向的数据, 其他都不能修改
slice 是结构体和指针的混合体
 */
type tt struct {
	id int
}
func main() {
	t := tt{123}
	p := &t
	fmt.Println(1, t, &p)
	change(t)
	p = &t
	fmt.Println(2, t, &p)
	t = change(t)
	p = &t
	fmt.Println(3, t, &p)
	changeP(p)
	p = &t
	fmt.Println(4, t, &p)

	for i := 0; i < 10; i++ {
		//引用
		defer func() {
			fmt.Println("func:i=", i)
		}()
		//值
		defer func(i int) {
			fmt.Println("func:i=", i)
		}(i)

		defer fmt.Println(i)
	}
}
//函数参数传值, 闭包传引用!
func change(t tt) tt {
	p := &t
	fmt.Println(5, t, &p)
	t.id += 100
	return t
}

func changeP(t *tt) {
	fmt.Println(6, t, &t)
	t.id += 200
}

