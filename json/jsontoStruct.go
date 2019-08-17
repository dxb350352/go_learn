package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	AAA string `json:"a"`
}
type TestB struct {
	BBB string `json:"b"`
}

func main() {
	var test Test
	var testb TestB
	m := map[string]string{"a":"1", "b":"2"}
	byt, _ := json.Marshal(m)
	json.Unmarshal(byt, &test)
	fmt.Println(test)
	json.Unmarshal(byt, &testb)
	fmt.Println(testb)
}
