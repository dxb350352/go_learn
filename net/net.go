package main
import (
	"fmt"
	"net"
	"bufio"
	"io"
	"time"
)
func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}
func main() {
	go startServer("9001")
	go startServer("9002")
	go startServer("9003")
	time.Sleep(time.Hour)
}

func startServer(port string) {
	ln, err := net.Listen("tcp", ":" + port)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go Echo(conn)
	}
}

func Echo(c net.Conn) {
	defer c.Close()
	for {
		line, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("Failure to read:%s\n", err.Error())
			break
		}
		fmt.Println(line)
//		_, err = c.Write([]byte(line))
//		if err != nil {
//			fmt.Printf("Failure to write: %s\n", err.Error())
//			break
//		}
	}
	fmt.Println("end......................")
}