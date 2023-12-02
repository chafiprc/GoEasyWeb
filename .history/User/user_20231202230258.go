package User

type UserInfo struct {
	UserID   uint32 `form:"userid" json:"user_id" gorm:"primaryKey"`
	Username string `form:"username" json:"user_name"`
	Pssword  string `form:"userpassword" json:"user_password"`
}

var IsLogin bool = false
var CurrentUser UserInfo