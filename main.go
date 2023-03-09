package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	age  int8
}

func FormatAsDate(t time.Time) string {
	yy, mm, dd := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", yy, mm, dd)
}

func main() {
	e := gee.Default()

	e.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})
	e.GET("/panic", func(c *gee.Context) {
		names := []string{"emo", "maitian"}
		c.String(http.StatusOK, names[100]) // ERROR: Index out of range
	})

	e.Run(":9999")
}

func middlewareV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
