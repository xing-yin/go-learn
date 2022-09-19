package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 参考：https://geektutu.com/post/gee-day1.html

// Engine is the uni handler for all requests
type Engine struct {
}

func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found:%s\n", req.URL)
	}
}

func main() {
	e := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", e))
}

// 封装前处理 JSON 的做法
func (e Engine) originDisposeJSON(w http.ResponseWriter, req *http.Request) {
	obj := map[string]interface{}{
		"username": "Bob",
		"password": "123",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(obj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// 封装后处理 JSON 的做法
func (e Engine) encapsulationDisposeJSON(w http.ResponseWriter, req *http.Request) {
	//c := context.Background()
	//c.JSON(http.StatusOK, gee.H{
	//	"username": c.PostForm("username"),
	//	"password": c.PostForm("password"),
	//})
}
