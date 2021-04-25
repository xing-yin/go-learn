package main

import (
	"fmt"
	"net/http"

	"time"
)

// 内置 http server 实现
func main() {
	// 定义不同路径下不同处理逻辑
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	// 它采⽤用最⻓长匹配原则，如果有多个匹配，一定采⽤用匹配路路径最⻓长的那个进⾏行行处 理理
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
