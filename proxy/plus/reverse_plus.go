package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var agency = []string{
	"webapi.amap.com",
	"vdata.amap.com",
	"webrd01.is.autonavi.com",
	"webrd02.is.autonavi.com",
	"webrd03.is.autonavi.com",
	"webrd04.is.autonavi.com",
	"restapi.amap.com",
	"gaode.com",
	"m.amap.com",
	"wprd01.is.autonavi.com",
	"wprd02.is.autonavi.com",
	"wprd03.is.autonavi.com",
	"wprd04.is.autonavi.com",
	"webst01.is.autonavi.com",
	"webst02.is.autonavi.com",
	"webst03.is.autonavi.com",
	"webst04.is.autonavi.com",
	"vector.amap.com",
	"grid.amap.com",
	"tm.amap.com",
	"yuntuapi.amap.com",
}

var timeout = 10
//监听端口的所有请求转发到配置的url
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	for port, v := range agency {
		go Listen(fmt.Sprint(":", 10000+port), v+":80")
	}
	wg.Wait()
}

func Listen(srcHost, agencyHost string) {
	if srcHost == "" || agencyHost == "" {
		fmt.Fprintf(os.Stderr, "Fatal error: %s,%s,%s", "参数为空", srcHost, agencyHost)
		os.Exit(1)
	}
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", srcHost)
	CheckError(err)
	defer netListen.Close()

	Log(srcHost, "Waiting for clients")
	for {
		conn, err := netListen.Accept()
		//如果没有请求就一直等待
		if err != nil {
			continue
		}
		if conn != nil {
			Log(conn.RemoteAddr().String(), " tcp connect success")
			go handleConnection(conn, agencyHost) //go 可以实现异步并发请求
		}
	}
}

//处理连接
func handleConnection(conn net.Conn, agencyHost string) {
	defer conn.Close()
	buffer := ReceiveData(conn)
	if len(buffer) > 1 {
		arr := strings.Split(string(buffer), "\n")
		if len(arr) > 1 {
			arr[1] = "Host: " + agencyHost
			newstr := strings.Join(arr, "\n")
			SendAgencyHost([]byte(newstr), agencyHost, conn)
		}
	}

}
func SendAgencyHost(data []byte, host string, baseconn net.Conn) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err, host)
		return
	}
	defer conn.Close()
	conn.Write(data)
	//进行转发
	buffer := ReceiveData(conn)
	//返回数据里的url，可以利用上面的监听配置替换
	baseconn.Write(buffer)
	//进行转发--这个方式会等待很他时间
	//go io.Copy(conn, baseconn)
	//io.Copy(baseconn, conn)
}

//接收数据统一方法
func ReceiveData(conn net.Conn) []byte {
	var buf bytes.Buffer
	buffer := make([]byte, 8192)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100))
		sizenew, err := conn.Read(buffer)
		if err == io.EOF || sizenew == 0 {
			break
		}
		buf.Write(buffer[:sizenew])
	}
	return buf.Bytes()
}

//打印信息统一方法
func Log(v ...interface{}) {
	log.Println(v...)
}

//执行错误处理方法
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func copyBuffer(dst io.Writer, src io.Reader, flag int) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	//if wt, ok := src.(io.WriterTo); ok {
	//	return wt.WriteTo(dst)
	//}
	//// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	//if rt, ok := dst.(io.ReaderFrom); ok {
	//	return rt.ReadFrom(src)
	//}
	var buf = make([]byte, 32*1024)
	var body []byte
	for {
		nr, er := src.Read(buf)
		fmt.Println(nr, er, flag)
		if nr > 0 {
			body = append(body, buf...)
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
