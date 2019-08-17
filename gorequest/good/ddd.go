package main

import (
	"github.com/sas/utils"
	"fmt"
)

func main() {
	//m := map[string]string{"area":"1"}
	//data, err := utils.ExecGetMethodUrl("https://192.168.130.149:9005/bmalerts/pullbyarea", m)
	//fmt.Println(data, err)
	rui:="https://192.168.130.240:9007/za1bbix/GetLastClcsByNames?log_source=3&host_names[]=zmkvirtual1"
	request, err := utils.GetSkipValidateRequest(rui)
	if err != nil {
		fmt.Println(err)
	}
	//查询
	_, body, errs := request.Get(rui).EndBytes()
	fmt.Println(string(body),errs)
}
