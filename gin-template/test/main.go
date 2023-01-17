package main

import (
	"fmt"
	"gintemp/utils"
)

func main() {
	resp := utils.NewResponse()
	utils.Default(resp)
	fmt.Println(*resp)
}
