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
		return
	} else {
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		id := findUserId(username)
		var nowUser User.UserInfo
		nowUser = findUser(id)
		if id==0{
			c.String(http.StatusOK,"You need to register first!")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/register")
			return
		} else {
			if password == nowUser.Pssword {
				User.IsLogin = true
				User.CurrentUser = nowUser
			} else {
				c.String(http.StatusOK,"密码错误，请重新输入")
				time.Sleep(3*time.Second)
				c.Redirect(http.StatusFound,"/Login")
				return
			}
		}
	}
}

func Register(c *gin.Context) {
	if User.IsLogin {
		c.String(http.StatusOK,"You have been logined in!")
		c.Redirect(http.StatusFound,"/")
		return
	} else {
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		repeatPwd := c.PostForm("repeatpassword")
		if password != repeatPwd {
			c.String(http.StatusOK,"不匹配的密码！请重新输入")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/Register")
			return
		}
		id := findUserId(username)
		if id != 0 {
			c.String(http.StatusOK,"该用户已经注册，请更换你的用户名")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/Register")
			return
		} else {
			var user User.UserInfo
			user.Username = username
			user.Pssword = password
			db.Create(&user)
		}
	}
}

func findUserId(username string) uint32 {
	var user User.UserInfo
	err := db.Where("user_name=?",username).First(&user).Error
	if err != nil {
		return 0
	}
	return user.UserID
}

func findUser(id uint32) (User.UserInfo) {
	var user User.UserInfo
	err := db.Where("user_id=?",id).First(&user).Error
	if err != nil {
		return nil
	}
	return user
}