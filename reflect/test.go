package main
import (
	"reflect"
	"fmt"
)


type Parent_ struct {
	Id   int
	Name string
}

type Child_ struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var p Parent_
	p.Id = 1
	p.Name = "pn"
	var c Child_
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)
	cv := reflect.ValueOf(c)

	for i := 0; i < pt.NumField(); i++ {
		cv.FieldByName(pt.Field(i).Name).Set(pv.FieldByName(pt.Field(i).Name))
	}
	fmt.Println(c)
}