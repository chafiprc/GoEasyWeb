package User

type UserInfo struct {
	UserID   uint32 `json:"user_id" gorm:"primaryKey"`
	Username string `json:"user_name"`
	Pssword  string `json:"user_password"`
}

var IsLogin bool = false
var CurrentUser UserInfo