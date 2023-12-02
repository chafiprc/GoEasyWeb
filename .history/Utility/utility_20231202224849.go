package Utility

import (
	"EasyWeb/User"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:chafiprc@tcp(Chafi-MateBook-E:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	db.AutoMigrate(&User.UserInfo{})
}

func Login(c *gin.Context) {
	if User.IsLogin {
		c.String(http.StatusOK,"You have been logined in!")
		c.Redirect(http.StatusFound,"/")
	} else {
		c.String(http.StatusOK,"Login TODO")
		var json Login
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.String(http.StatusBadRequest,"Error!")
		} else {
			
		}
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

func findUser(username string) int {
	db.
}