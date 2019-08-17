package main

import (
	"time"
	"fmt"
	"github.com/jackiedong168/gorequest"
	"github.com/kr/pretty"
)

func main() {
	url:=`http://192.168.130.201:9000/search/search?keywords=__index=gaxz* __notimefilter=true  person_email=aa OR person_email="kriscross23@yahoo.aca"&from=0&size=10`
	request := gorequest.New().Timeout(time.Minute).Post(url)
	request.Param("", "")
	_, body, errs := request.EndBytes()
	pretty.Println(string(body))
	fmt.Println(errs)

}
