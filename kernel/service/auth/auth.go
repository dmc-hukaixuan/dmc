package auth

import (
	"dmc/initialize/database"
	"dmc/kernel/model/user"
	"dmc/kernel/util"
	"fmt"
)

//
func UserAuth(u *user.LoginParam) (err error, userEnter *user.User) {
	var user user.User
	// user password sha256 encryption
	u.PW = util.SHA2(u.PW)
	// .Preload("Authorities").Preload("Authority")
	// preload can be find users role, this is
	err = database.Gorm().Table("users").Where("login = ? AND pw = ?", u.Login, u.PW).First(&user).Error
	fmt.Println("user", userEnter, " err:", err)
	return err, &user
}
