package main

import (
	"path/filepath"
	"github.com/sas/gaxz/app/models"
	"fmt"
)

func main() {
	gogo("e:/fd/fds.txt")
	gogo("e:/fd/fds.cvs")
	gogo("e:/fd/fds.txt.retyr1")
	gogo("e:/fd/fds.txt.retyr2")
	gogo("e:/fd/fds.txt.retyr3")
	gogo("e:/fd/fds.txt.failed")

}

func gogo(filePath string) {
	dir, filename := filepath.Split(filePath)
	fname, extend := models.GetFileNameExtend(filename)
	switch extend {
	///!!!
	case "cvs", "txt":
		filePath += ".retry1"
	case "retyr1":
		filePath = filepath.Join(dir, fname + ".retry2")
	case "retyr2":
		filePath = filepath.Join(dir, fname + ".retry3")
	case "retyr3":
		filePath = filepath.Join(dir, fname + ".failed")
	}
	fmt.Println(filePath)
}
