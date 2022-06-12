package user

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"

	"fmt"
	"strconv"
)

func DepartmentList(validID int) map[string]string {
	var tp []model.Deparment
	selectSQL := `SELECT id, name FROM dmc_department`
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

func DepartmentGet(district_id int) (department model.Deparment) {
	err := global.GVA_DB.Table("dmc_department").Where("id = ?", district_id).First(&department).Error
	if err != nil {
		return
	}
	return department
}

func DepartmentAdd(department model.Deparment) (department_id int, err error) {
	err = global.GVA_DB.Table("dmc_department").Create(&department).Error
	if err != nil {
		return
	}
	return department.ID, err
}

func DepartmentUpdate(department model.Deparment) (district_id int, err error) {
	fmt.Println("add department -----------", department)
	err = global.GVA_DB.Table("dmc_department").Where("id = ?", department.ID).Model(&department).Omit("create_by", "create_time").Updates(department).Error
	fmt.Println("add department -----------", err)
	if err != nil {
		panic(err)
	}

	return department.ID, err
}

/*
	type list get
*/
func DepartmentListGet(validID int) (ts []model.Deparment) {
	selectSQL := `SELECT dd.id, dd.name, dd.valid_id, dd.street AS street,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					dd.create_time AS create_time, dd.change_time AS change_time
					FROM dmc_department dd 
					LEFT JOIN users u ON u.id = dd.create_by 
					LEFT JOIN users u1 ON u1.id = dd.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&ts).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return ts
}
