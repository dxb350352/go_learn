package main
import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"testgo/socket/wouq"
)
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "www.baidu.com:80")
	wouq.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	wouq.CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	wouq.CheckError(err)
	result, err := ioutil.ReadAll(conn)
	wouq.CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}