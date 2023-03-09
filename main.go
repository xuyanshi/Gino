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

	e.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := e.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=emo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := e.Group("/v2")
	v2.AddMiddleware(func(c *gee.Context) {
		t := time.Now()
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	})

	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/emo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		// curl "http://localhost:9999/v2/login" -X POST -d "username=emo&password=emo123456"
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	err := e.Run(":9999")
	if err != nil {
		return
	}
	// Useless
	// cmd := exec.Command("curl http://localhost:9999/v2/hello/emo\n")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()
}
