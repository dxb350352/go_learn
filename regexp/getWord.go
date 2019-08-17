package main

import (
	"regexp"
	"fmt"
	"strings"
)

func main() {
	getapptype()
}

func getapptype()  {
	str :=" __app=wondersoft_sndlp AND __type=t_eve "
	var reg = regexp.MustCompile(`((__app=)|(__type=))\S+`)
	r := reg.FindAllStringSubmatch(str, -1)
	for _, v := range r {
		if strings.HasPrefix(v[0], "__app=") {
			fmt.Println("app",v[0][6:])
		} else if strings.HasPrefix(v[0], "__type=") {
			fmt.Println("type",v[0][7:])
		}
	}
}

func getword()  {
	str :="f b aa     d-    c  e   a  d "
	var reg = regexp.MustCompile(`\S+`)
	r := reg.FindAllStringSubmatch(str, -1)
	for _, v := range r {
		fmt.Println(v)
	}
}

func gethostnamee()  {
	str:=`<189>2016/11/22 14:13:55 SW-45-1LQ-H-F1-1333 %%01SEC/5/STREAM(l): 过去5分钟会话连接数统计: 正常关闭的连接数 = 1， 关闭的TCP半连接数 = 188， 策略阻断的连接数 = 0， 查路由失败的连接数 = 0， 黑名单阻断的连接数 = 0。`
	reg:=regexp.MustCompile(` (F|S)W(-[0-9a-zA-Z]+)+ `)
	r := reg.FindAllStringSubmatch(str, -1)
	for _, v := range r {
		fmt.Println(v[0])
	}
}