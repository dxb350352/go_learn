package main

import (
	"github.com/cihub/seelog"
	"io/ioutil"
	"strings"
	"os"
	"fmt"
)

var Logger seelog.LoggerInterface

func main() {
	byt, err := ioutil.ReadFile("E:/GOPATH/src/testgo/log/config/log.xml")
	if err != nil {
		panic(err)
	}
	testConfig := string(byt)
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(host)
	testConfig = strings.Replace(testConfig, "{hostname}", host, 1)
	Logger, err = seelog.LoggerFromParamConfigAsString(testConfig, nil)
	if err != nil {
		panic(err)
	}
	defer Logger.Flush()
	Logger.Error("logger started............")
}
