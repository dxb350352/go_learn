package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	v := url.Values{}
	v.Set("dsn", "jdbc:informix-sqli://192.168.130.160:9090/db_test:")
	v.Set("driver", "com.informix.jdbc.IfxDriver")
	v.Set("user", "AAO")
	v.Set("password", "123456")
	v.Set("sql", "SELECT * FROM systables")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	//body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://192.168.130.181:8080/tm/query", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	fmt.Printf("%+v\n", req)                                                         //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}