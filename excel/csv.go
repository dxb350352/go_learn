package main

import (
	"encoding/csv"
	"os"
	"log"
	"io"
	"fmt"
)

func main() {
	csvReader,err:=os.Open(`E:\GOPATH\src\testgo\excel\wx.csv`)
	if err!=nil{
		log.Fatal(err)
	}
	defer csvReader.Close()
	csvR:=csv.NewReader(csvReader)
	for{
		arr,err:=csvR.Read()
		if err==io.EOF{
			break
		}else if err!=nil{
			continue
		}
		fmt.Println(arr)
	}

}
