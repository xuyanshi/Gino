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

func TestGetRoute2(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/g/o", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)

	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatalf("pattern should be /assets/*filepath and filepath should be file1.txt.\n"+
			"but now pattern is %s and filepath is %s.", n1.pattern, ps1["filepath"])
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	if n2 != nil && ps2 != nil {
		ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
		if !ok2 {
			t.Fatalf("pattern should be /assets/*filepath & filepath should be css/test.css.\n"+
				"but now pattern is %s and filepath is %s.", n1.pattern, ps1["filepath"])
		}
	}
	t.Fatal("r.getRoute(\"GET\", \"/assets/css/test.css\") returns nil")
}

func TestGetRoutes(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/g/o", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filename", nil)

	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes should be 4")
	}
}
