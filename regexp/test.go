package main
import (
	"regexp"
	"fmt"
)

func main() {
	domain := "http://x0.wondersoft.com:9005"
	var reg = regexp.MustCompile(`https*://\d+\.\d+\.\d+\.\d+`)
	s := "<a href=http://192.168.130.240:9005/users/signin/abc>"
	r := reg.ReplaceAllString(s, domain+"?redirect=${0}")
	fmt.Println(r)


	reg=regexp.MustCompile(`^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$`)
	fmt.Println(reg.MatchString(`120521198610090312`))
}
