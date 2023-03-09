package gee

import "fmt"

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
			}
		}()

		c.Next()
	}
}
