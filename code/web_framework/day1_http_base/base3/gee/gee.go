package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handlerFunc
}

// GET defines the method to add GET request
func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRouter("GET", pattern, handlerFunc)
}

// POST defines the method to add POST request
func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRouter("POST", pattern, handlerFunc)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}
