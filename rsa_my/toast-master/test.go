package main

import (
	"fmt"
	"encoding/base64"
	"testgo/rsa_my/toast-master/rsa"
	"testgo/rsa_my/toast-master/crypto"
	crypto1 "crypto"
)

func main() {
	plant := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	key, err := rsa.LoadKeyFromPEMFile(
		`E:/GOPATH/src/testgo/rsa_my/toast-master/crypto/rsa_public_key.pem`,
		`E:/GOPATH/src/testgo/rsa_my/toast-master/crypto/rsa_private_key.pem`,
		rsa.ParsePKCS8Key)
	if err != nil {
		fmt.Println(err)
		return
	}

	cipher, err := crypto.NewRSA(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	enT, err := cipher.Encrypt([]byte(plant))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(enT))

	deT, err := cipher.Decrypt(enT)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(deT))

	signBytes, err := cipher.Sign([]byte(plant), crypto1.SHA1)
	if err != nil {
		fmt.Println(err)
		return
	}

	sign := base64.StdEncoding.EncodeToString(signBytes)

	fmt.Println(sign)

	errV := cipher.Verify([]byte(plant), signBytes, crypto1.SHA1)
	if errV != nil {
		fmt.Println(errV)
		return
	}
}
