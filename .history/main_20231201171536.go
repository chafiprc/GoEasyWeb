package main

import (
	"WebModule/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Router.RouterInit(r)
}