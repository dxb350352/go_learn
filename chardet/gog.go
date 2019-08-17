package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"bufio"
	"io"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/hydra13142/chardet"
)

func main() {
	path1 := "d:/003-003-2-1-10.txt"
	path2 := "d:/003-002-2-1-10.txt"
	printFileContentchardet(path1)
	printFileContentchardet(path2)

}

func printFileContent(path string) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(byt))
}
func printFileContentchardet2(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal()
	}
	r := bufio.NewReader(file)
	byt := make([]byte, 10)
	a, b := r.Read(byt)
	fmt.Println(a, b, string(byt))
}

func NewReaderByEncodeDet(file *os.File) io.Reader {
	byt := make([]byte, 10240)
	count, err := file.Read(byt)
	file.Seek(0,0)
	//fmt.Println(count,string(byt))
	if err != nil {
		return file
	}
	if count > 0 {
		var reader io.Reader
		encode := chardet.Mostlike(byt)
		switch encode {
		case "utf-8":
			reader = file
		case "big5", "gbk":
			reader = transform.NewReader(file, simplifiedchinese.GBK.NewDecoder())
		case "gb18030":
			reader = transform.NewReader(file, simplifiedchinese.GB18030.NewDecoder())
		}
		return reader
	}
	return file
}

func printFileContentchardet(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal()
	}
	reader := NewReaderByEncodeDet(file)
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}