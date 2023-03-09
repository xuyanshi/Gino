package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func main() {
	e := gee.New()
	e.AddMiddleware(gee.Logger())
	e.Static("/assets", "./static")

	e.Run(":9999")
}

func middlewareV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
