package main

import (
	"fmt"
	"net/http"
	"testing"
)

// http.ResponseWriter：代表响应，传递到前端的
// *http.Request：表示请求，从前端传递过来的
func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1 style='color:red'>hello Golang!<h1>")
}

func TestHttp(b *testing.T) {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v \n", err)
		return
	}
}
