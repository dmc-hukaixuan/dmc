package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

func DistrictList(validID int) map[string]string {
	var tp []model.TicketPriority
	selectSQL := `SELECT id, name FROM district`
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

func DistrictGet(district_id int) (district model.District) {
	err := global.GVA_DB.Table("district").Where("id = ?", district_id).First(&district).Error
	if err != nil {
		return
	}
	return district
}

func DistrictAdd(district model.District) (district_id int, err error) {
	err = global.GVA_DB.Table("district").Create(&district).Error
	if err != nil {
		return
	}
	return district.ID, err
}

func DistrictUpdate(district model.District) (district_id int, err error) {
	err = global.GVA_DB.Table("ticket_priority").Where("id = ?", district.ID).Model(&district_id).Omit("create_by", "create_time").Updates(district).Error
	if err != nil {
		return
	}
	return district.ID, err
}

/*
	type list get
*/
func DistrictListGet(validID int) (ts []model.TicketSource) {
	selectSQL := `SELECT tt.id, tt.name, tt.valid_id, tt.tnstart AS tnstart,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					tt.create_time AS create_time, tt.change_time AS change_time
					FROM district tt 
					LEFT JOIN users u ON u.id = tt.create_by 
					LEFT JOIN users u1 ON u1.id = tt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&ts).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return ts
}
