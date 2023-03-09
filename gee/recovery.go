package gee

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {

		}()

		c.Next()
	}
}
