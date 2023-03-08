package gee

import (
	"fmt"
	"reflect"
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

func TestParsePattern(t *testing.T) {
	if !(reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"}) &&
		reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"}) &&
		reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})) {
		t.Fatal("TestParsePattern Failed.")
	}
}
