package main

import (
	"fmt"
	"regexp"
)
//护照，香港，澳门，台湾，大陆
var Exp_person_id = regexp.MustCompile(`(^1[45][0-9]{7}|([P|p|S|s]\d{7})|([S|s|G|g]\d{8})|([Gg|Tt|Ss|Ll|Qq|Dd|Aa|Ff]\d{8})|([H|h|M|m]\d{8,10})$)|(^((\s?[A-Za-z])|([A-Za-z]{2}))\d{6}(([0−9aA])|(\([0-9aA]\)))$)|(^[1|5|7][0-9]{6}\([0-9Aa]\))|(^[a-zA-Z][0-9]{9}$)|(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)`)

func main() {
	fmt.Println("护照", Exp_person_id.MatchString("g28233515"))
	fmt.Println("香港", Exp_person_id.MatchString("x354670(0)"))
	fmt.Println("香港", Exp_person_id.MatchString("b354670a"))
	fmt.Println("香港", Exp_person_id.MatchString("g354670a"))
	fmt.Println("澳门", Exp_person_id.MatchString("5215299(8)"))
	fmt.Println("台湾", Exp_person_id.MatchString("a123456789"))
	fmt.Println("台湾", Exp_person_id.MatchString("e223963669"))
	fmt.Println("台湾", Exp_person_id.MatchString("y243418722"))
	fmt.Println("台湾", Exp_person_id.MatchString("e191713473"))
	fmt.Println("台湾", Exp_person_id.MatchString("v290828926"))
	fmt.Println("台湾", Exp_person_id.MatchString("b258545149"))
	fmt.Println("大陆", Exp_person_id.MatchString("320882198902162412"))
	fmt.Println("大陆", Exp_person_id.MatchString("120105199201018916"))
	fmt.Println("大陆", Exp_person_id.MatchString("120105199201019556"))
	fmt.Println("大陆", Exp_person_id.MatchString("350725199001012772"))
	fmt.Println("大陆", Exp_person_id.MatchString("350725199001012158"))
}