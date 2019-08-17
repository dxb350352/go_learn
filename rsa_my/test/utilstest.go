package main

import (
	rsa_my "github.com/sas/utils"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/base64"
)

func main() {
	rsa_my.SetKeyFileBasePath("E:/GOPATH/src/github.com/sas/sas")
	//createRnFile()
	readRnFile()
	//text()
	//sign()
	//verify()
}

func text() {
	str := "fddfdfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdsfsa"
	str += "fddfdfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdsfsa"
	sasrn, err := rsa_my.Encrypt([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("...........")
	sasrn, err = rsa_my.Decrypt(sasrn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(sasrn))
}
func createRnFile() {
	sasrn, err := ioutil.ReadFile("E:/GOPATH/src/testgo/rsa_my/test/sas.rn")
	if err != nil {
		fmt.Println(err)
		return
	}
	sasrn, err = rsa_my.Encrypt(sasrn)
	if err != nil {
		fmt.Println(err)
		return
	}
	ofile, err := os.Create("E:/GOPATH/src/testgo/rsa_my/test/sas_r.rn")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ofile.Close()
	str := base64.StdEncoding.EncodeToString(sasrn)
	_, err = ofile.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func readRnFile() {
	sasrn, err := ioutil.ReadFile("E:/GOPATH/src/testgo/rsa_my/test/sas_r.rn")
	if err != nil {
		fmt.Println(err)
		return
	}
	sasrn, err = base64.StdEncoding.DecodeString(string(sasrn))
	if err != nil {
		fmt.Println(err)
		return
	}
	sasrn, err = rsa_my.Decrypt(sasrn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(sasrn), "........")
}

func sign() {
	sasrn, err := ioutil.ReadFile("E:/GOPATH/src/testgo/rsa_my/test/sas.rn")
	if err != nil {
		fmt.Println(err)
		return
	}

	sasrn, err = rsa_my.Sign(sasrn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%x", sasrn)
}

func verify()  {
	sasrn, err := ioutil.ReadFile("E:/GOPATH/src/testgo/rsa_my/test/sas.rn")
	if err != nil {
		fmt.Println(err)
		return
	}

	signrn, err := rsa_my.Sign(sasrn)
	if err != nil {
		fmt.Println(err)
		return
	}
	bl:=rsa_my.Verify(sasrn,signrn)
	fmt.Println(bl)
}
