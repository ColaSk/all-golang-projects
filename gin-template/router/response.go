package router

import (
	"github.com/gin-gonic/gin"
)

/*
status
data
message
*/

type response struct {
	status  int
	code    int
	data    any
	message string `default:"success"`
}

func NewResponse() *response {

	return &response{}

}

func Response(c *gin.Context, status int, data gin.H) {
	c.JSON(status, data)
}
