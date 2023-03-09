package main

import (
	"fmt"
	"gee"
	"html/template"
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
	e := gee.New()
	e.AddMiddleware(gee.Logger())
	e.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./static")

	stu1 := &student{
		Name: "emo",
		age:  23,
	}
	stu2 := &student{
		Name: "maitian",
		age:  24,
	}
	e.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	e.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":   "gee",
			"student": [2]*student{stu1, stu2},
		})
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
