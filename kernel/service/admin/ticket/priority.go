/*
All ticket priority functions.
*/
package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
	get ticket priority

	return a map data stucture
*/
func PriorityList(validID int) map[string]string {
	var tp []model.TicketPriority
	selectSQL := `SELECT id, name FROM ticket_priority`
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

/*
	get priority detail
*/
func PriorityGet(priorityID int) (tp model.TicketPriority) {
	err := global.GVA_DB.Table("ticket_priority").Where("id = ?", priorityID).First(&tp).Error
	if err != nil {
		return
	}
	return tp
}

/* ticket priority */
func PriorityAdd(tp model.TicketPriority) (prioritID int, err error) {
	err = global.GVA_DB.Table("ticket_state").Create(&tp).Error
	if err != nil {
		return
	}
	return tp.ID, err
}

/*
  ticket priority id
*/
func PriorityUpdate(tp model.TicketPriority) (prioritID int, err error) {
	err = global.GVA_DB.Table("ticket_priority").Where("id = ?", tp.ID).Model(&tp).Omit("create_by", "create_time").Updates(tp).Error
	if err != nil {
		return
	}
	return tp.ID, err
}

/*
	Priority list get
*/
func PriorityListGet(validID int) (tp []model.TicketPriority) {
	selectSQL := `SELECT tt.id, tt.name, tt.valid_id, tt.tnstart AS tnstart,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					tt.create_time AS create_time, tt.change_time AS change_time
					FROM ticket_source tt 
					LEFT JOIN users u ON u.id = tt.create_by 
					LEFT JOIN users u1 ON u1.id = tt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return tp
}
