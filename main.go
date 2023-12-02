package main

import (
	"EasyWeb/Router"
	"EasyWeb/Utility"

	"github.com/gin-gonic/gin"
)

func main() {
	Utility.InitDB()
	r := gin.Default()
	Router.RouterInit(r)
	r.Run(":8080")
}