/*
   All user function.
*/
package user

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

func UserList(validID int) map[string]string {
	var tp []model.TicketPriority
	selectSQL := `SELECT id, name FROM users`
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
		PriorityList[strconv.Itoa(v.ID)] = v.Name
	}
	return PriorityList
}

func UserGet() {

}

func UserAdd() {

}

func UserUpdate() {

}
