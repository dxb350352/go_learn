package main

import (
	"fmt"
	"encoding/base32"
	"time"
	"math/rand"
	"github.com/skip2/go-qrcode"
)

func main() {
	key := GetRandomString(10)
	key = base32.StdEncoding.EncodeToString([]byte(key))
	fmt.Println(key)
	str := "otpauth://totp/booing?secret=" + key
	err := qrcode.WriteFile(str, qrcode.Medium, 256, "d:/sss.png")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetRandomString(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
