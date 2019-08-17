package main

import (
	"path/filepath"
	"fmt"
	"os"
)

func main() {
	pat := "d:/revel.exe"
	dir, name := filepath.Split(pat)
	dir = filepath.Dir(dir)
	dir += "_finish"
	os.MkdirAll(dir, 666)
	pat2 := filepath.Join(dir, name)
	err := os.Rename(pat, pat2)
	fmt.Println(err)
}
