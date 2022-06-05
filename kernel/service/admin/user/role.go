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

func RoleList(validID int) map[string]string {
	var tp []model.Role
	selectSQL := `SELECT id, name FROM role`
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
		fmt.Println("v", v.ID)
		PriorityList[strconv.Itoa(v.ID)] = v.Name
	}
	return PriorityList
}

func RoleGet(roleID int) (role model.Role) {

	// selectSQL := `SELECT id, name, web, mobile, describes, valid_id, icon, color, type,
	// 			 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
	err := global.GVA_DB.Table("queue").Where("id = ?", roleID).First(&role).Error
	if err != nil {
		return
	}
	return role
}

func RoleAdd() {

}

func RoleUpdate() {

}
