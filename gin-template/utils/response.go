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

func (r *response) Default() {
	Default(r)
}

func NewResponse() *response {
	r := &response{}
	r.Default()
	return r
}

func Response(c *gin.Context, status int, code int, data response) {
	c.JSON(status, data)
}
