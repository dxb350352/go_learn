package rsa_my

import (
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"crypto/rand"
	"encoding/pem"
	"strings"
	"errors"
	"bytes"
)

const privateKeyFile = "E:/GOPATH/src/testgo/rsa_my/toast-master/crypto/rsa_private_key.pem"
const publicKeyFile = "E:/GOPATH/src/testgo/rsa_my/toast-master/crypto/rsa_public_key.pem"

func ReadPemFile(file string) ([]byte, error) {
	file = strings.TrimSpace(file)

	pukBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return make([]byte, 0), err
	}

	puk, _ := pem.Decode(pukBytes)
	if puk == nil {
		return make([]byte, 0), errors.New("publicKey is not pem formate")
	}
	return puk.Bytes, nil
}

func GetPrivateKeyPKCS8() (*rsa.PrivateKey, error) {
	privateKey, err := ReadPemFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	prkI, err := x509.ParsePKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return prkI.(*rsa.PrivateKey), nil
}

func GetPublicKey() (*rsa.PublicKey, error) {
	publicKey, err := ReadPemFile(publicKeyFile)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}
//加密
func Encrypt(inputstr []byte) ([]byte, error) {
	public, err := GetPublicKey()
	if err != nil {
		return make([]byte, 0), err
	}
	groups := grouping(inputstr, len(public.N.Bytes()) - 11)
	buffer := bytes.Buffer{}
	for _, plainTextBlock := range groups {
		cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, public, plainTextBlock)
		if err != nil {
			return nil, err
		}
		buffer.Write(cipherText)
	}
	return buffer.Bytes(), nil
	//return rsa.EncryptPKCS1v15(rand.Reader, public, inputstr)

}
//解密
func Decrypt(inputstr []byte) ([]byte, error) {
	private, err := GetPrivateKeyPKCS8()
	if err != nil {
		return make([]byte, 0), err
	}
	groups := grouping(inputstr, len(private.N.Bytes()))
	buffer := bytes.Buffer{}
	for _, plainTextBlock := range groups {
		cipherText, err := rsa.DecryptPKCS1v15(rand.Reader, private, plainTextBlock)
		if err != nil {
			return nil, err
		}
		buffer.Write(cipherText)
	}
	return buffer.Bytes(), nil
	//return rsa.DecryptPKCS1v15(rand.Reader, private, inputstr)
}

/*数据太长的时候，要按照秘钥的长度对数据进行分组*/
func grouping(src []byte, size int) [][]byte {
	var groups [][]byte
	srcSize := len(src)
	if srcSize <= size {
		groups = append(groups, src)
	} else {
		for len(src) != 0 {
			if len(src) <= size {
				groups = append(groups, src)
				break
			} else {
				v := src[:size]
				groups = append(groups, v)
				src = src[size:]
			}
		}
	}
	return groups
}