package main

import (
	"regexp"
	"fmt"
)

func main() {
	expfomat := regexp.MustCompile(`^\[\s*([\p{Han}\w\._-]+)\s*\]\s*((EXCLUDE)|(INCLUDE))\s*([\p{Han}\w\._-]+)\s*$`)
	fmt.Println(expfomat.MatchString("[fdsa地] INCLUDE dd地"))
	//身份证
	expfomat = regexp.MustCompile(`^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$`)
	fmt.Println("身份证:", expfomat.MatchString(""))
	//邮箱
	expfomat = regexp.MustCompile(`^([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+(\.[a-zA-Z]{2,3})+$`)
	fmt.Println("邮箱:", expfomat.MatchString("dd@d.com"))
	//手机
	expfomat = regexp.MustCompile(`^1[3|4|5|8]\d{9}$`)
	fmt.Println("手机:", expfomat.MatchString("13540064782"))
	//固话
	expfomat = regexp.MustCompile(`^(0\d{2,3}-)?([2-9]\d{6,7})+(-\d{1,4})?$`)
	fmt.Println("固话:", expfomat.MatchString("9057559461"))
	//QQ
	expfomat = regexp.MustCompile(`^\d{5,10}$`)
	fmt.Println("QQ:", expfomat.MatchString("658974587"))
	//银行卡
	expfomat = regexp.MustCompile(`^(\d{16}|\d{19})$`)
	fmt.Println("银行卡:", expfomat.MatchString("44444444414444336444"))
}
