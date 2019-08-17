package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8081", "E:/GOPATH_GO/src/testgo/gohttps/2-https/server.crt", "E:/GOPATH_GO/src/testgo/gohttps/2-https/server.key", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
