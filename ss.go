package main
import (
	"github.com/Knetic/govaluate"
	"fmt"
)

func main() {
	//1
	//		expression, _ := govaluate.NewEvaluableExpression("10 > 0")
	//		result,_ := expression.Evaluate(nil)
	//		fmt.Println(result)
	//2
	//		expression, _ := govaluate.NewEvaluableExpression("foo > 0")
	//		parameters := make(map[string]interface{}, 8)
	//		parameters["foo"] = -1
	//		result, _ := expression.Evaluate(parameters)
	//		fmt.Println(result)
	//3
	//		expression, _ := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	//		parameters := make(map[string]interface{}, 8)
	//		parameters["requests_made"] = 100
	//		parameters["requests_succeeded"] = 90
	//		result, _ := expression.Evaluate(parameters)
	//		fmt.Println(result)
	//4
	//		expression, _ := govaluate.NewEvaluableExpression("http_response_body == 'service is ok'")
	//		parameters := make(map[string]interface{}, 8)
	//		parameters["http_response_body"] = "service is ok"
	//		result, _ := expression.Evaluate(parameters)
	//		fmt.Println(result)
	//5
	//	expression, _ := govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	//	parameters := make(map[string]interface{}, 8)
	//	parameters["mem_used"] = 512
	//	parameters["total_mem"] = 1024
	//	result, _ := expression.Evaluate(parameters)
	//	fmt.Println(result)
	//6
	expression, _ := govaluate.NewEvaluableExpression("[c] > 2 ||  [a] > 2 || [b] > 1")
	parameters := make(map[string]interface{}, 8)
	var x float64 = 2
	parameters["a"] = x
	parameters["b"] = x
	parameters["c"] = x
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
