package Utility

import (
	"EasyWeb/User"
	"net/http"
	"time"

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
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		id := findUser(username)
		if id==0{
			c.String(http.StatusOK,"You need to register first!")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/register")
		} else {
			if password == findPwd(id) {
				User.IsLogin = 
			}
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

func findUser(username string) uint32 {
	var user User.UserInfo
	err := db.Where("user_name=?",username).First(&user).Error
	if err != nil {
		return 0
	}
	return user.UserID
}

func findUser(id uint32) string {
	var user User.UserInfo
	err := db.Where("user_id=?",id).First(&user).Error
	if err != nil {
		panic("获取密码时发生错误！")
	}
	return user.Pssword
}