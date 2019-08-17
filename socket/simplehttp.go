package main
import (
	"net"
	"os"
	"fmt"
	"testgo/socket/wouq"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	wouq.CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	wouq.CheckError(err)
	result, err := wouq.ReadFully(conn)
	wouq.CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}