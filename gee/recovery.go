package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// PC is the program counter for the location in this frame.
// For a frame that calls another frame, this will be the
// program counter of a call instruction. Because of inlining,
// multiple frames may have the same PC value, but different
// symbolic information.
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	return ""
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", message)
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
