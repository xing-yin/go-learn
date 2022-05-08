package main

import (
	"fmt"
	"net/http"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome,Gopher\n")
}

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(greeting))
	//http.ListenAndServe(":8088", greeting())
}
