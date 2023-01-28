package main

import (
	"gone"
	"log"
)

func main() {

	engine := gone.New()

	engine.POST("/hello", helloHandler)

	log.Fatal(engine.Run(":9999"))

	// http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Header
func helloHandler(c *gone.Context) {
	c.WriteJSON(200, map[string]interface{}{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
