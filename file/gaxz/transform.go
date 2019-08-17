package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
	"io/ioutil"
	"time"
)

func main() {
	//os.Args = append(os.Args, "d:/testff")
	//os.Args = append(os.Args, "E:/GOPATH/src/testgo/file/gaxz")
	if len(os.Args) < 2 {
		fmt.Println("请输入目标文件夹")
		return
	}
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	outdir := filepath.Join(os.Args[1], "transform")
	os.MkdirAll(outdir, os.ModeDir)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "txt") || strings.HasSuffix(file.Name(), "csv") {
			input := filepath.Join(os.Args[1], file.Name())
			output := filepath.Join(outdir, file.Name())
			start := time.Now().Unix()
			fmt.Println("开始处理文件：", input)
			transform(input, output)
			end := time.Now().Unix()
			fmt.Println("结束处理文件：", input, "，用时：", end - start, "秒")
		}
	}
}

func transform(input, output string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	infile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer infile.Close()
	outfile, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)
	reader := bufio.NewScanner(infile)
	reader.Scan()
	line := reader.Text()
	header := strings.Count(line, "|")
	writer.WriteString(line + "\n")
	line = ""
	var count int
	for reader.Scan() {
		count++
		text := strings.TrimSpace(reader.Text())
		if strings.Count(text, "|") < header {
			line += text
			if strings.Count(line, "|") < header {
				continue
			}
		} else {
			line = text
		}

		writer.WriteString(line + "\n")
		line = ""
		if count % 10000 == 0 {
			writer.Flush()
			fmt.Println("flush lines......10000,handled:", count)
		}
	}
	writer.Flush()
	fmt.Println("flush last lines......,handled:", count)
}
