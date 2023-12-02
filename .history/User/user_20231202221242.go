package User

type User struct {
	Username string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword  string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

var IsLogin bool = false
var CurrentUser User