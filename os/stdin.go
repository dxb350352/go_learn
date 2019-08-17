package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//
	var i, j, k int
	fmt.Scanln(&i, &j, &k)
	fmt.Println(i, j, k)
	//
	byts := make([]byte, 100)
	os.Stdin.Read(byts)
	fmt.Println(string(byts))
	//
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader.ReadLine())

	ScanLine()
}

func ScanLine() {
	var c byte
	var err error
	var b []byte
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	fmt.Println(string(b))
}
