package main

import (
	"os"
	"github.com/op/go-logging"
	"fmt"
)

func main() {

	fileName := "d://log.txt"
	logFile, err := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err, ".........")
	}
	log := logging.MustGetLogger("ss")
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled)

	channelbk:=logging.NewChannelMemoryBackend(10)
	channelbackend1Leveled:=logging.AddModuleLevel(channelbk)
	channelbackend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled,channelbackend1Leveled)

	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("xiaorui.cc")
	log.Critical("太严重了")

	arr := []interface{}{"1", "2", "3"}
	log.Error(arr...)
}