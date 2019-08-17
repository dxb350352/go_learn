package main

import (
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Gender bool
}

var user User

func main() {
	u:=new(User)
	fmt.Println(u)
	fmt.Println(user)
	fmt.Println(user.Age)
	//值传递
	changeuser(user)
	fmt.Println(user)
}
func changeuser(u User) {
	u.Name = "ddd"
}
