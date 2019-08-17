package ca

import (
	"io/ioutil"
	"crypto/x509"
	"errors"
	"crypto/tls"
)

func PopulateCertPool(cartPath string) (*x509.CertPool, error) {
	pem, err := ioutil.ReadFile(cartPath)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pem) {
		return nil, errors.New("Failed appending certs")
	}
	return certPool, nil
}


func CerficateFromPEMs(crt, key string) (tls.Certificate, error) {
	mycert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	return mycert, nil
}
