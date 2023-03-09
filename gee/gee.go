package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(c *Context)

type (
	RouterGroup struct {
		prefix      string
		middleWares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
	Engine struct {
		routerGroup *RouterGroup
		router      *router
		groups      []*RouterGroup // store all groups here
	}
)

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	// 在 Go 语言中，实现了接口方法的 struct 都可以强制转换为接口类型
	return http.ListenAndServe(addr, engine)

	// 也可以这样写
	/*
		handler := (http.Handler)(engine) // 强制转换为 http.Handler 类型
		log.Fatal(http.ListenAndServe(":9999", handler))
	*/
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
