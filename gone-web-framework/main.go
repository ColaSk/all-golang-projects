package main

import (
	"gone"
	"net/http"
)

func main() {
	r := gone.New()
	r.GET("/", func(c *gone.Context) {
		c.WriteHTML(http.StatusOK, "<h1>Hello gone</h1>")
	})

	r.GET("/hello", func(c *gone.Context) {
		// expect /hello?name=gonektutu
		c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gone.Context) {
		// expect /hello/gonektutu
		c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gone.Context) {
		c.WriteJSON(http.StatusOK, map[string]interface{}{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")
}
