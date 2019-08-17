package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"regexp"
)

func main() {
	var str string = "[c] >= 4 || ([a] >= 2 || [b] <= 1) "
	fmt.Println(str)
	arr := GetTriggerParameters(str)
	fmt.Println(arr)
	parameters := make(map[string]interface{})
	var v float64 = 2
	for _, key := range arr {
		parameters[key] = v
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
