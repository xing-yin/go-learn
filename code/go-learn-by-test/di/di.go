package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello,%s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "http")
}

func main() {
	// 通过引入 io.Writer 接口，让 Greet 函数变得更有通用性
	Greet(os.Stdout, "Elen")

	// http.ResponseWriter 也实现了 io.Writer 接口
	http.ListenAndServe(":8000", http.HandlerFunc(MyGreetHandler))
}
