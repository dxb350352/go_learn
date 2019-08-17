package main
import (
	"fmt"
	"regexp"
)

func main() {
	loginRegex := regexp.MustCompile("^(\\w|\\.|-){1,15}$")
	fmt.Println(loginRegex.MatchString("b_11fdsafdsfadd_11.-.."))
	loginRegex = regexp.MustCompile("^/user/.*$")
	fmt.Println(loginRegex.MatchString("/user/rew/rereww"))

	fmt.Println(regexp.Match("\\[\\d*\\]", []byte("paths[0]")))

//	fmt.Println(regexp.MatchString("\\*|^([1-5]?[0-9]{1})$|^([1-5]?[0-9]{1})$-^([1-5]?[0-9]{1})$)|^([1-5]?[0-9]{1})$/^([1-5]?[0-9]{1})$)", "59"))

	fmt.Println(regexp.MatchString(`[\\u4e00-\\u9fa5]+`,"中文"))
	var reg = regexp.MustCompile(`\[\s*(\w)+\s*\]`)
	r := reg.FindAllStringSubmatch("[中文] > d", -1)
	for _, v := range r {
		fmt.Println(v)
	}
}