package main
import (
	"time"
	"fmt"
)

type name struct {
	t time.Time
}
func main() {
	var n name
	fmt.Println(n.t)
	n.t = time.Now()
	fmt.Println(n.t)
	n.t = *new(time.Time)
	fmt.Println(n.t)

}