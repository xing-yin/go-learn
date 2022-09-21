package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup // store all groups
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRouter(method string, comp string, handlerFunc HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Router %4s - %s", method, pattern)
	g.engine.router.addRouter(method, pattern, handlerFunc)
}

// GET defines the method to add GET request
func (g *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	g.addRouter("GET", pattern, handlerFunc)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	g.addRouter("POST", pattern, handlerFunc)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}

// ServeHTTP 当接收到一个具体请求时，要判断该请求适用于哪些中间件，这里我们简单通过 URL 的前缀来判断。得到中间件列表后，赋值给 c.handlers。
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middleware []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middleware = append(middleware, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middleware
	e.router.handle(c)
}
