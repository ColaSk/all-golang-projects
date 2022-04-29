package main

import (
	"gee"
	"net/http"
)

// func onlyForV2() gee.HandlerFunc {
// 	return func(c *gee.Context) {
// 		// Start timer
// 		t := time.Now()
// 		// if a server error occurred
// 		c.Fail(500, "Internal Server Error")
// 		// Calculate resolution time
// 		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
// 	}
// }

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
