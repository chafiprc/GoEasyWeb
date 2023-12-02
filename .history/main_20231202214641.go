package main

import (
	"EasyWeb/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Router.RouterInit(r)
	r.Run(":8080")
}