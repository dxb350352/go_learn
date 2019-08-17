package main
import (
	"fmt"
	"time"
	"math/rand"
	"strings"
)

func main() {
	s := fmt.Sprint("aa", "bb", "cc", time.Now().Unix(), rand.Intn(999999))
	fmt.Println(s)
	s = fmt.Sprint("aa", "bb", "cc", time.Now().Unix(), "ddd", rand.Intn(999999))
	fmt.Println(s)
	s = fmt.Sprint("aa", "bb", "cc", time.Now().Unix(), "", rand.Intn(999999))
	fmt.Println(s)
	s = ":9090"
	fmt.Println(s[strings.LastIndex(s, ":")+1:])
	fmt.Println("123\r456")

	fmt.Println("1234567890"[:6])
	var a=1
	var b=3
	fmt.Println(a/b)
}
