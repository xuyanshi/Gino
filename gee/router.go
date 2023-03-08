package gee

import (
	"log"
	"net/http"
)

type router struct {
	roots    map[string]*node       // eg. roots["GET"], roots["POST"]
	handlers map[string]HandlerFunc // eg. handlers["GET-/p/:lang/doc"], handlers["POST-/p/book"]
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
