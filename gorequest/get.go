package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	for i := 0; i < 20000; i++ {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://192.168.130.181:8080/tm/query?dsn=jdbc:informix-sqli://192.168.130.160:9090/db_test:&driver=com.informix.jdbc.IfxDriver&user=AAO&password=123456&sql=select%20*%20from%20systables", ioutil.NopCloser(strings.NewReader("")))

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
		fmt.Printf("%+v\n", req)                                                         //看下发送的结构

		resp, err := client.Do(req) //发送
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(i, "-------", string(data), err)
		resp.Body.Close()     //一定要关闭resp.Body
	}
}