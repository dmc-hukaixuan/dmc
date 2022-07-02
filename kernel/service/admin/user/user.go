/*
   All user function.
*/
package user

import (
	"dmc/global"
	model "dmc/kernel/model/user"
	"fmt"
	"strconv"
)

func UserList(validID int) map[string]string {
	var tp []model.User
	selectSQL := `SELECT id, full_name, job_number, login FROM users`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	PriorityList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return PriorityList
	}

	for _, v := range tp {
		PriorityList[strconv.Itoa(v.ID)] = v.FullName + " " + v.JobNumber + " " + v.Login
	}
	return PriorityList
}

func UserGet(userID int) (userData model.User) {

	return userData
}

func UserAdd() {

}

func UserUpdate() {

}
