package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func beforeRequest(c *gin.Context) {
	fmt.Println("request start ...")
	method := c.Request.Method
	url := c.Request.RequestURI
	fmt.Printf("method: %s, URI: %s", method, url)
}

func afterRequest(c *gin.Context) {
	fmt.Println("request end ...")
}

func defaultMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		beforeRequest(c)
		c.Next()
		afterRequest(c)
	}
}

func corsMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Next()
	}
}

func InitMiddleware(engine *gin.Engine) {

	engine.Use(corsMiddle(), defaultMiddle())

}
