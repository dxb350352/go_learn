package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.TLS.PeerCertificates, ".................perr")
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "E:/GOPATH_GO/src/testgo/gohttps/6-dual-verify-certs/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.VerifyClientCertIfGiven,
		},
	}

	err = s.ListenAndServeTLS("E:/GOPATH_GO/src/testgo/gohttps/6-dual-verify-certs/server.crt", "E:/GOPATH_GO/src/testgo/gohttps/6-dual-verify-certs/server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
