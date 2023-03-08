package gee

import (
	"fmt"
	"testing"
)

func TestGetRouter(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/g/o", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filename", nil)

	n, ps := r.getRoute("GET", "/hello/golang")
	if n == nil {
		t.Fatal("nil shouldn't be returned.")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("matching error!")
	}

	if ps["name"] != "golang" {
		t.Fatal("name error")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}
