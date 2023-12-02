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
	dsn := "root:chafiprc@tcp(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
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
		nowUser := findUser(id)
		if id==0{
			c.AbortWithStatus(200)
			c.Writer.WriteString("You need to register first!\nRedirecting...")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/register")
			return
		} else {
			if password == nowUser.Pssword {
				User.IsLogin = true
				User.CurrentUser = nowUser
			} else {
				c.AbortWithStatus(200)
				c.Writer.WriteString("密码错误，请重新输入\nRedirecting...")
				time.Sleep(3*time.Second)
				c.Redirect(http.StatusFound,"/login")
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
			c.AbortWithStatus(200)
        	c.Writer.WriteString("不匹配的密码！请重新输入\nRedirecting...")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/register")
			return
		}
		id := findUserId(username)
		if id != 0 {
			c.AbortWithStatus(200)
        	c.Writer.WriteString("该用户已经注册，请更换你的用户名\nRedirecting...")
			time.Sleep(3*time.Second)
			c.Redirect(http.StatusFound,"/register")
			return
		} else {
			var user User.UserInfo
			user.Username = username
			user.Pssword = password
			db.Create(&user)
			User.IsLogin = true
			User.CurrentUser = user
		}
	}
}

func findUserId(username string) uint32 {
	var user User.UserInfo
	err := db.Where("username=?",username).First(&user).Error
	if err != nil {
		return 0
	}
	return user.UserID
}

func findUser(id uint32) (User.UserInfo) {
	var user User.UserInfo
	err := db.Where("user_id=?",id).First(&user).Error
	if err != nil {
		return User.UserInfo{}
	}
	return user
}