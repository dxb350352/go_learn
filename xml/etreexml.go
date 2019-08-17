package main

import (
	"github.com/beevik/etree"
	"fmt"
	"math/rand"
	"encoding/base64"
	"crypto/md5"
	"io"
	"github.com/sas/utils"
)

func main() {
	utils.SetKeyFileBasePath("E:/GOPATH/src/github.com/sas/sas")
	doc := etree.NewDocument()
	common := doc.CreateElement("common")
	common.CreateElement("deadline").CreateCharData("2014-3-3")
	common.CreateElement("sa-key").CreateCharData("sfdsafdsafds")
	moduleList := doc.CreateElement("moduleList")
	module := moduleList.CreateElement("module")
	module.CreateElement("ID").CreateCharData("12")
	module.CreateElement("Number").CreateCharData(fmt.Sprint(rand.Intn(4)))
	module.CreateElement("Version").CreateCharData("705")
	module = moduleList.CreateElement("module")
	module.CreateElement("ID").CreateCharData("13")
	module.CreateElement("Number").CreateCharData(fmt.Sprint(rand.Intn(4)))
	module.CreateElement("Version").CreateCharData("705")
	//doc.WriteTo(os.Stdout)
	//生成XML文件内容
	str, err := doc.WriteToString()
	if err != nil {
		fmt.Println(err.Error())
	}
	//加密
	rsastr, err := utils.Encrypt([]byte(str))
	if err != nil {
		fmt.Println(err.Error())
	}
	//base64
	base64str := base64.StdEncoding.EncodeToString(rsastr)
	//md5
	h := md5.New()
	io.WriteString(h, base64str)

	md := fmt.Sprintf("%X", h.Sum(nil))
	//最终文件内容
	str = md + str
	fmt.Println(str)
}
