package main
import "fmt"

type NilTest struct {
	Name string
}

func (n NilTest) test() {
	fmt.Println("test", n)
}

func (n *NilTest) test2() {
	fmt.Println("test2")
}

func main() {
	var n NilTest
	n.test()
	n.test2()
	fmt.Println(".............................")
	var n2 *NilTest
	n2.test()
	n2.test2()
}
