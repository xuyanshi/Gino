package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			return
		}
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}
	})

	r.Run(":9999")
}
