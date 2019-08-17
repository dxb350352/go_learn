package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"archive/zip"
	"io"
	"compress/gzip"
	"archive/tar"
	"path/filepath"
	"strings"
	"github.com/sas/utils"
)

func main() {
	testzip()
	//testgzip()
	//path := "C:/Users/Administrator/Desktop/log样本整理/windowslog/Logs"
	//DeCompressGzip(path + "/test.tar.gz", path + "/sasslvpn_web")
}

func testzip() {
	path := "D:/Desktop/layer-v2.4"
	file := "D:/Desktop/layer-v2.4.zip"
	err := CompressZip(path, file)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DeCompressZip(file, path + "/zip")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func testgzip() {
	path := "C:/Users/Administrator/Desktop/log样本整理/windowslog/Logs"
	file := "C:/Users/Administrator/Desktop/log样本整理/windowslog/Logs/gzip.tar.gz"
	err := CompressGzip(path + "/DPX", file)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DeCompressGzip(file, path + "/gzip")
	if err != nil {
		fmt.Println(err.Error())
	}
}
//压缩ZIP
func CompressZip(path, fileName string) error {
	fzip, err := os.Create(fileName)
	if err != nil {
		return err
	}
	w := zip.NewWriter(fzip)
	defer w.Close()
	return CompressZipFolder(w, path, "")
}
//压缩ZIP递归
func CompressZipFolder(w  *zip.Writer, path string, folder string) error {
	path = filepath.Clean(path)
	if len(folder) == 0 {
		folder = "/"
	}
	//获取源文件列表
	info, err := os.Stat(filepath.Clean(path + folder))
	if err != nil {
		return err
	}
	var f []os.FileInfo
	if info.IsDir() {
		f, err = ioutil.ReadDir(filepath.Clean(path + folder))
		if err != nil {
			return err
		}
	} else {
		f = append(f, info)
	}
	for _, file := range f {
		if file.IsDir() {
			err = CompressZipFolder(w, path, folder + file.Name() + "/")
			if err != nil {
				return err
			}
			continue
		}
		fw, err := w.Create(folder + file.Name())
		if err != nil {
			return err
		}
		filecontent, err := ioutil.ReadFile(path + folder + file.Name())
		if err != nil {
			return err
		}
		_, err = fw.Write(filecontent)
		if err != nil {
			return err
		}
	}
	return nil
}
//解压ZIP
func DeCompressZip(fileName, dir string) error {
	dir = filepath.Clean(dir)
	//删除目录及以下文件
	if utils.CheckFileExist(dir) {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(dir, 0777) //创建一个目录
	if err != nil {
		return err
	}

	cf, err := zip.OpenReader(fileName) //读取zip文件
	if err != nil {
		return err
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		fmt.Println(file.Name, file.FileInfo().IsDir())
		separator := "/"
		if strings.HasPrefix(file.Name, "/") {
			separator = ""

		}
		path := filepath.Dir(dir + separator + file.Name)
		if !utils.CheckFileExist(path) {
			err = os.MkdirAll(path, 0777)//创建一个目录
			if err != nil {
				return err
			}
		}
		f, err := os.Create(dir + separator + file.Name)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
//压缩GZIP
func CompressGzip(path, fileName string) error {
	// file write
	fw, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	return CompressGzipFolder(tw, path, "")
}
//压缩GZIP递归
func CompressGzipFolder(tw *tar.Writer, path, folder string) error {
	path = filepath.Clean(path)
	if len(folder) == 0 {
		folder = "/"
	}
	//获取源文件列表
	info, err := os.Stat(filepath.Clean(path + folder))
	if err != nil {
		return err
	}
	var f []os.FileInfo
	if info.IsDir() {
		f, err = ioutil.ReadDir(filepath.Clean(path + folder))
		if err != nil {
			return err
		}
	} else {
		f = append(f, info)
	}

	for _, fi := range f {
		if fi.IsDir() {
			err = CompressGzipFolder(tw, path, folder + fi.Name() + "/")
			if err != nil {
				return err
			}
			continue
		}
		// 打开文件
		fr, err := os.Open(path + folder + fi.Name())
		if err != nil {
			return err
		}
		defer fr.Close()
		// 信息头
		h := new(tar.Header)
		h.Name = folder + fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()
		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			return err
		}
		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			return err
		}
	}
	return nil
}
//解压GZIP
func DeCompressGzip(fileName, dir string) error {
	dir = filepath.Clean(dir)
	//删除目录及以下文件
	if utils.CheckFileExist(dir) {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(dir, 0777) //创建一个目录
	if err != nil {
		return err
	}

	// file read
	fr, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(h.Name, h.FileInfo().IsDir())
		separator := "/"
		if strings.HasPrefix(h.Name, "/") {
			separator = ""

		}
		path := filepath.Dir(dir + separator + h.Name)
		if !utils.CheckFileExist(path) {
			err = os.MkdirAll(path, 0777)//创建一个目录
			if err != nil {
				return err
			}
		}
		f, err := os.Create(dir + separator + h.Name)
		if err != nil {
			return err
		}
		defer f.Close()
		// 写文件
		_, err = io.Copy(f, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

