package Router

import "github.com/gin-gonic/gin"

func RouterInit(r *gin.Engine) {
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}