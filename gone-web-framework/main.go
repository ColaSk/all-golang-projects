package main

import (
	"gone"
	"net/http"
)

func main() {
	r := gone.New()
	r.GET("/index", func(c *gone.Context) {
		c.WriteHTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gone.Context) {
			c.WriteHTML(http.StatusOK, "<h1>Hello gone</h1>")
		})

		v1.GET("/hello", func(c *gone.Context) {
			// expect /hello?name=gonektutu
			c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gone.Context) {
			// expect /hello/gonektutu
			c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gone.Context) {
			c.WriteJSON(http.StatusOK, map[string]string{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
