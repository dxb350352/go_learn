package main

import (
	"github.com/jackiedong168/gocaptcha"
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("d://test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	id := gocaptcha.NewLen(4)
	gocaptcha.WriteImageCeil(file, id, 200, 100,100)
}
