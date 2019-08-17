package main

import (
	"runtime"
	"fmt"
)

type Vector []int

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u[i]
	}
	c <- 1
}

func (v Vector) DoAll(u Vector) {
	cpu := runtime.NumCPU()
	fmt.Println("CPU:", cpu)
	c := make(chan int, cpu)
	for i := 0; i < cpu; i++ {
		go v.DoSome(i*len(u)/cpu, (i+1)*len(u)/cpu, u, c)
	}
	for i := 0; i < cpu; i++ {
		<-c
	}
	fmt.Println(v)
}

func main() {
	v := Vector([]int{1, 2, 3, 4})
	v.DoAll(v)
}
