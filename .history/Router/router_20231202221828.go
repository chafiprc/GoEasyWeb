package Router

import (
	"EasyWeb/User"
	"EasyWeb/Utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	r.LoadHTMLGlob("HTML/*")
	r.GET("/", func(c *gin.Context) {
		var content string
		if !User.IsLogin {
			content = "Sorry, you need to login or register."
		} else {
			content = "Hi, I am glad to see you, " + User.CurrentUser.Username + "!"
		}
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"EasyWebsite!",
			"content":content,
		})
	})
	r.GET("/login",func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{})
	})
	r.GET("/register",func(c *gin.Context) {
		c.HTML(http.StatusOK,"register.html",gin.H{})
	})
	r.POST("/loginRequest",Utility.Login)
}