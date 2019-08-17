package main

import (
	"github.com/Knetic/govaluate"
	"fmt"
	"regexp"
	"strings"
	"github.com/bitly/go-simplejson"
)

func main() {
	js, _ := simplejson.NewJson([]byte(`{"fds":"1","aaa":2,"c":true}`))
	fmt.Println(GetTriggerParameters1(`fds={fds}&fdsddd={aaa}&c={c}`, js))
	var str string = "[a]>=  2   &&   [b]<=2 "
	fmt.Println(str)
	arr := GetTriggerParameters(str)
	fmt.Println(arr)
	parameters := make(map[string]interface{})
	for _, key := range arr {
		parameters[key] = 1
	}
	fmt.Println(parameters)
	expression, err := govaluate.NewEvaluableExpression(str)
	if err != nil {
		fmt.Println(err)
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func init() {
	fmt.Println(".....................")
}
func GetTriggerParameters(exp string) (result []string) {
	var reg = regexp.MustCompile(`\[\s*(\w+)\s*\]`)
	r := reg.FindAllStringSubmatch(exp, -1)
	for _, v := range r {
		fmt.Print("----")
		fmt.Println(v)
		//第二个哦
		result = append(result, v[1])
	}
	return
}
func GetTriggerParameters1(exp string, js *simplejson.Json) string {
	var reg = regexp.MustCompile(`\{\s*(\w+)\s*\}`)
	r := reg.FindAllStringSubmatch(exp, -1)
	for _, v := range r {
		nstr := fmt.Sprint(js.Get(v[1]))
		nstr = nstr[2:len(nstr) - 1]
		exp = strings.Replace(exp, v[0], nstr, -1)
	}
	return exp
}