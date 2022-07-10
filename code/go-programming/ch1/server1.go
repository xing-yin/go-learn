package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	//http.HandleFunc("/", Handler)
	//http.HandleFunc("/", HandlerV2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// do something
	})
	http.HandleFunc("/count", counterHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func counterHandler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(writer, "Count %d\n", count)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func HandlerV2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAdd = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
