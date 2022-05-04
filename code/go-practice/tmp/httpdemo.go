package main

import (
	"log"
	"net/http"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/3
 * @Desc:
 */

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello", hello)
	log.Println("Starting http server ...")
	log.Fatal(http.ListenAndServe(":50052", nil))
}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello World"))
}

func pong(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("pong"))
}
