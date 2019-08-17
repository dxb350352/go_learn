package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	barr,err:=ioutil.ReadFile("E:/GOPATH/src/testgo/file/insert/casetype.txt")
	if err!=nil{
		fmt.Println(err)
	}
	items:=strings.Split(string(barr),"\n")
	for i,v:=range items{
		item:=strings.Split(strings.TrimSpace(v),"\t")
		if len(item)>=2{
			fmt.Printf("INSERT INTO `tqry`.`sys_codes` ( `code`, `name`, `parent_code`, `pid`, `sort_no`, `enable`, `create_time`, `update_time`) VALUES ( '%s', '%s', '%s', '%d', '11', '1', '2016-10-26 15:51:37', '2016-10-26 15:51:37');\n",item[0],item[1],"CaseType",i+1)
		}
	}
}
