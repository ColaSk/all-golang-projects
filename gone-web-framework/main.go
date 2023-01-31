package main

import (
	"gone"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gone.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gone.Context) {
			// expect /hello?name=gonektutu
			c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")

	v2.Use(func(c *gone.Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	})
	{
		v2.GET("/hello/:name", func(c *gone.Context) {
			// expect /hello/gonektutu
			log.Printf("zzz")
			c.WriteString(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gone.Context) {
			c.WriteJSON(http.StatusOK, map[string]string{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
		v2.GET("/panic", func(c *gone.Context) {
			names := []string{"geektutu"}
			c.WriteString(http.StatusOK, names[100])
		})
	}

	r.Run(":9999")
}
