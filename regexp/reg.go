package main

import (
	"fmt"
	"regexp"
)

func main() {
	loginRegex := regexp.MustCompile("^[a-zA-Z]([a-zA-Z0-9]|_){5,20}$")
	fmt.Println(loginRegex.MatchString("b_11_11"))

	loginRegex = regexp.MustCompile("^__app=.*;$")
	fmt.Println(loginRegex.MatchString("__app=fdsfds;"))

	a, b, c := 1, 2, 3
	a = b + c
	b = a + c
	c = a + b
	fmt.Println(a, b, c)

	phone := regexp.MustCompile("(^[0-9]{3,4}\\-[0-9]{3,8}$)|(^[0-9]{7,8}$)|(^\\([0-9]{3,4}\\)[0-9]{3,8}$)|(^0{0,1}1[0-9]{10}$)")
	fmt.Println(phone.MatchString("1234567"))
}
