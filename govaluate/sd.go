package main

import (
	"github.com/Knetic/govaluate"
	"fmt"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("[非法]")
	if err != nil {
		fmt.Print(err)
	}
	result, err := expression.Evaluate(map[string]interface{}{"非法":true})
	fmt.Println(result,err)
}
