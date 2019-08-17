package main

import (
	"fmt"
)

type notifier interface {
	Notify()
}

type User struct {
	Name string
	Age  int
}

func (s *User) Notify() {
	fmt.Println("name is: ", s.Name)
}

func notifination(n notifier) {
	n.Notify()
}
//https://www.jb51.net/article/56831.htm
func main() {
	//编译通过
	u := &User{"james", 33}
	//编译不通过
	//u := User{"james", 33}
	notifination(u)
}
