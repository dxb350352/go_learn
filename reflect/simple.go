package main
import (
	"reflect"
	"fmt"
)

func main() {
	var i string
	i = "fdsa"
	gogo(i)
}

func gogo(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t.Name())
}