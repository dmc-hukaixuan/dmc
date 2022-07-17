/*
   All user function.
*/
package user

import (
	"dmc/global"
	model_template "dmc/kernel/model/ticket"
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

// get user role lsit
func UserRoleList(userID int) (roles []int) {
	var userRole []model.UserRole
	selectSQL := `SELECT id, queue_id, user_id FROM queue_user`
	err := global.GVA_DB.Raw(selectSQL, userID).Scan(&userRole).Error
	if err != nil {
		panic(err)
	}
	for _, v := range userRole {
		roles = append(roles, v.QueueID)
	}
	return roles
}

func UserCreateTemplateList(userID int) map[int]model_template.TemplateData {
	var template []model_template.TemplateData
	sql := `SELECT tt.id, tt.name FROM dmc_ticket_template tt
				LEFT JOIN dmc_ticket_template_role ttr ON ttr.template_id = tt.id
			WHERE tt.type = 'create'
					AND ( ttr.role_id = 0 OR ttr.role_id IN ( SELECT queue_user.queue_id FROM queue_user where queue_user.user_id = ?))`
	global.GVA_DB.Raw(sql, userID).Model(&template)

	templateList := map[int]model_template.TemplateData{}
	//
	for _, v := range template {
		templateList[v.ID] = v
	}
	return templateList
}
