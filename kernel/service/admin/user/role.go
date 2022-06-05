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
	selectSQL := `SELECT id, name FROM queue`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	roleList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return roleList
	}
	for _, v := range tp {
		roleList[strconv.Itoa(v.ID)] = v.Name
	}
	return roleList
}

func RoleOverview() (role []model.Role) {
	selectSQL := `SELECT q.id as id, q.name AS name, q.default_owner AS default_owner, q.default_responsible AS default_responsible,
					q.calendar_name AS calendar_name, q.system_address_id AS system_address_id, q.comments AS comments, 
					u.full_name AS create_by_name, u1.full_name AS change_by_name, q.create_time AS create_time, q.change_time AS change_time
					FROM queue q LEFT JOIN users u ON u.id = q.create_by LEFT JOIN users u1 ON u1.id = q.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&role).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return role
}

func RoleGet(roleID int) (role model.Role) {
	// selectSQL := `SELECT id, name, web, mobile, describes, valid_id, icon, color, type,
	// 			 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
	err := global.GVA_DB.Table("queue").Where("id = ?", roleID).First(&role).Error
	if err != nil {
		return
	}
	fmt.Println(" role get  ", role)
	return role
}

func RoleAdd(roleData model.Role) (roleID int, err error) {
	err = global.GVA_DB.Table("queue").Create(&roleData).Error
	if err != nil {
		return
	}
	return roleData.ID, err
}

/*
	role update
*/
func RoleUpdate(roleData model.Role) (roleID int, err error) {
	err = global.GVA_DB.Table("queue").Where("id = ?", roleData.ID).Model(&roleData).Omit("create_by", "create_time").Updates(roleData).Error
	if err != nil {
		return
	}
	return roleData.ID, err
}
