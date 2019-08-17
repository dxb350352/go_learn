package main
import "fmt"

type Proxy struct {
	Name string
	Url  string
	Ip   string
}

type GetProxy func() [] Proxy

var getProxy GetProxy

func InitFilter(proxy GetProxy) {
	getProxy = proxy
}

func main() {
	var myproxy = func() [] Proxy {
		var p []Proxy
		p = append(p, Proxy{Name:"1"})
		p = append(p, Proxy{Name:"11"})
		return p
	}
	InitFilter(myproxy)
	ss := [] Proxy(getProxy())
	for _, v := range ss {
		fmt.Println(v.Name)
	}
}