package main

import (
	"fmt"
	"net/http"
)

/// 自己实现一个简易的路由器
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello myroute")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9091", mux)
}
