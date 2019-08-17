package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	"strings"
)

func main() {
	u4 := uuid.NewV4()
	fmt.Println(strings.Replace(u4.String(), "-", "", -1))
	fmt.Println(len(strings.Replace(u4.String(), "-", "", -1)))
}
