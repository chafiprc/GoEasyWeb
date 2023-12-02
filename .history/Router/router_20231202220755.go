package Router

import (
	"net/http"
	"EasyWeb/User"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	r.LoadHTMLGlob("HTML")
	r.GET("/", func(c *gin.Context) {
		if User
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"EasyWebsite!",
			"content":
		})
	})
}