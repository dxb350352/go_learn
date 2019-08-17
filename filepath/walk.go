package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	path := "E:/log"
	files := make([]string, 0)
	//walk是个好方法啊，直接就遍历所有文件了,文件夹里的文件也要走一个

	filepath.Walk(path, func(pathfile string, info os.FileInfo, err error) error {
		if info.IsDir() {
			//说好的所有，文件夹也算，因此要剔除掉
			return nil
		}
		fmt.Println(pathfile,"               ", info.Name())
		files = append(files, pathfile[len(path):])
		return nil
	})
	fmt.Println(files)
}