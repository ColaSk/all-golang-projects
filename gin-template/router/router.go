package router

import (
	"fmt"
	"gintemp/middlewares"
	"gintemp/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	middlewares.InitMiddleware(router)

	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("pang")
		utils.Response(c, 200, 200, *utils.NewResponse())
	})

	router.Run()
}
