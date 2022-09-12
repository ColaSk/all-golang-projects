package router

import (
	"fmt"
	"gintemp/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	middlewares.InitMiddleware(router)

	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("pang")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
