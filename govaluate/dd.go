package main

import (
	"github.com/Knetic/govaluate"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/sas/utils"
)

func main() {
	exp, err := govaluate.NewEvaluableExpression("[总流量占用率] > 60")
	if err != nil {
		fmt.Println(err)
	}
	tj:=simplejson.New()
	tj.Set("network_used","100.0")
	result, err := exp.Evaluate(map[string]interface{}{"总流量占用率":utils.ParseFloat64(fmt.Sprint(tj.Get("network_used").Interface()))})
	//result, err := exp.Evaluate(map[string]interface{}{"总流量占用率":tj.Get("network_used").Interface()})
	flag, isbool := result.(bool)
	fmt.Println(result,flag,isbool)


}
