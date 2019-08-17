package main
import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("ip4:icmp", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn.LocalAddr(), conn.RemoteAddr())
}