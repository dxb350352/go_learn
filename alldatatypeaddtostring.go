package main
import (
	"fmt"
	"time"
)

func main() {
	var a int64 = 123
	var b float64 = 234.234
	var t time.Time = time.Now()
	str := fmt.Sprint(a, "+", b, "=", t)
	fmt.Println(str)
}
