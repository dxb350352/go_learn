package main

import "fmt"

type arrinterface interface {

}

type stringarr []string


func main() {
	printObject("printObject")
	stringarr:=[]string{"a", "b"}
	printArray(stringarr)
}

func printMap(key string, m map[string]interface{}) {
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("....................printMap")
}

func printArray(arr arrinterface) {
	arri:=arr.([]interface{})
	for i, v := range arri {
		fmt.Println(i, v)
	}
	fmt.Println("....................printArray")
}

func printObject(str interface{}) {
	fmt.Println(str)
	fmt.Println("....................printObject")
}