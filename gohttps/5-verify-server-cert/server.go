package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081",
		"E:/GOPATH_GO/src/testgo/gohttps/5-verify-server-cert/server.crt", "E:/GOPATH_GO/src/testgo/gohttps/5-verify-server-cert/server.key", nil)
}
