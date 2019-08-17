package main
import (
	"github.com/cihub/seelog"
	"fmt"
	"errors"
)

var Logger seelog.LoggerInterface

func init() {
	var err error
	Logger, err = seelog.LoggerFromConfigAsFile("./src/testgo/log/log3.xml")
	if err != nil {
		fmt.Println(err.Error(),"....err")
		panic(err)
	}
	Logger.Info("logger started............")
	Logger.Warn("logger started............")
	Logger.Error("logger started............")
	Logger.Debug("logger started............")
	Logger.Flush()
}


func main() {
	Logger.Info("gogogo")
	Logger.Errorf("err:%# v",errors.New("error test gogogogogo"))
	Logger.Flush()
	Logger.Info("gogogo")
	Logger.Errorf("err:%# v",errors.New("error test gogogogogo"))
	Logger.Flush()
}
