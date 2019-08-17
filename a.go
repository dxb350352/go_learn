package main

import (
	"fmt"
	"encoding/hex"
	"testgo/sm4"
)

func main() {
	key := []byte("1234567890abcdef")
	iv := []byte("1234567890abcdef")
	in := []byte("北京明朝万达科技有股份有限公司")
	fmt.Println(hex.EncodeToString(in))

	out := sm4.EncryptCbcPadding(key, iv, in)
	fmt.Println(hex.EncodeToString(out))
	out2 := sm4.DecryptCbcPadding(key, iv, out)
	fmt.Println(hex.EncodeToString(out2))
}
