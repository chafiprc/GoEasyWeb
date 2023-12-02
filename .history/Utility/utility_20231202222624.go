package Utility

import (
	"EasyWeb/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if User.IsLogin {
		c.String(http.StatusOK,"You have been logined in!")
		c.Redirect(http.StatusFound,"/")
	} else {
		c.String(http.StatusOK,"Login TODO")
	}
}

func Register(c *gin.Context) {
	if User.IsLogin {
		c.String(http.StatusOK,"You have been logined in!")
		c.Redirect(http.StatusFound,"/")
	} else {
		c.String(http.StatusOK,"Register TODO")
	}
}