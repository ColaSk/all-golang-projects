package utils

import (
	"github.com/gin-gonic/gin"
)

/*
status
data
message
*/

type response struct {
	Status  int `default:"200"`
	Code    int `default:"200"`
	Data    any
	Message string `default:"success"`
}

func NewResponse() *response {

	return &response{}

}

func Response(c *gin.Context, status int, data gin.H) {
	c.JSON(status, data)
}
