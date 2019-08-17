package main

import (
	"math/rand"
	"os"
	"fmt"
	"bufio"
)

func main() {
	dataFile, err := os.OpenFile("d:/gogo.svm", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dataFile.Close()
	wr := bufio.NewWriter(dataFile)
	for i := 0; i < 5000; i++ {
		t := random(0, 86400)
		label := 0
		if t <= 28800 {
			label = 1
		}
		fmt.Printf("%d 1:28672 2:28680 3:%d\n", label, t)
		wr.WriteString(fmt.Sprintf("%d 1:28672 2:28680 3:%d\n", label, t))
		if i % 100 == 99 {
			wr.Flush()
		}
	}
	wr.Flush()
}

func random(start, end int64) int64 {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//return r.Int63n(end - start) + start
	//rand.Seed(time.Now().UnixNano())
	return rand.Int63n(end - start) + start
}