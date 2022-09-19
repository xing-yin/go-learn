package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	e.router.addRouter(method, pattern, handlerFunc)
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
	// 调用 router.handle 之前，构造了一个 Context 对象
	context := newContext(w, req)
	e.router.handle(context)
}
