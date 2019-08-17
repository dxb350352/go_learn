package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println(time.Now().Format("20060102-150405"))

	t, err := time.Parse("Jan 02 15:04:05", "Nov 27 03:35:01")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format("20060102-150405"))
	t = t.AddDate(2016, 0, 0)
	fmt.Println(t.Format("20060102-150405"))
	fmt.Println( rand.Intn(9999))
	fmt.Println(time.Now().Format("20060102150405") + fmt.Sprintf("%04d", rand.Intn(9999)))
	fmt.Println( rand.Intn(9999))
	fmt.Println( rand.Intn(9999))
}
