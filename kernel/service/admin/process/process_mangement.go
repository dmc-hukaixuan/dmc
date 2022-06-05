package process

import (
	"dmc/global"
	datatime "dmc/kernel/util/time"
	"fmt"

	//"dmc/initialize/database"
	"dmc/kernel/model/admin"
	"time"
)

type ProcessModel struct{}

var ProcessModelA = new(ProcessModel)

// get process list
func ProcessTypeList() (list interface{}, total int64, err error) {
	var processType []admin.ProcessTypeList
	// make a sql string
	selectSQL := `SELECT id, name, valid_id, (SELECT COUNT(ppc.id) FROM pm_process_c ppc
				  WHERE ppc.process_type = ptc.id ) as count_process FROM  pm_process_type_c ptc`
	err = global.GVA_DB.Raw(selectSQL).Scan(&processType).Error
	if err != nil {
		return
	}
	return processType, total, err
}

type processList struct {
	ID               int       `json:"id" gorm:"column:id;"`
	Name             string    `json:"name" gorm:"column:name;"`
	StateEntityID    string    `json:"state_entity_id" `
	ProcessType      string    `json:"typename" gorm:"column:typename;"`
	CreateTime       time.Time `json:"create_time"`
	CreateByUserName string    `json:"create_by_name" gorm:"column:create_by_name;"`
	ChangeTime       time.Time `json:"change_time"`
	ChangeByUserName string    `json:"change_by_name" gorm:"column:change_by_name;"`
}

// get process list data
func ProcessList() (list interface{}, total int64, err error) {
	var process []processList

	// make a sql string
	selectSQL := `SELECT pp.id as id, pp.name as name, pp.state_entity_id, pp.create_time as create_time, pp.change_time as change_time,
				  u1.full_name as create_by_name, u2.full_name as change_by_name, ppc.name as typename
				  FROM pm_process_c pp LEFT JOIN users u1 ON u1.id = pp.create_by
				  LEFT JOIN pm_process_type_c ppc ON ppc.id = pp.process_type
				  LEFT JOIN users u2 ON u2.id = pp.change_by`
	err = global.GVA_DB.Raw(selectSQL).Scan(&process).Error

	// var results []map[string]interface{}
	// err = database.Gorm().Raw(selectSQL).Scan(&results).Error
	// fmt.Println("results LLL*------", results)

	if err != nil {
		return
	}
	return process, total, err
}

// process type detail get
func ProcessTypeGet(process_type_id int) (processType admin.ProcessType, err error) {
	err = global.GVA_DB.First(&processType, process_type_id).Error
	if err != nil {
		return
	}
	return processType, err
}

// process type detail get
func ProcessTypeAdd(pt *admin.ProcessType) (processTypeEnter admin.ProcessType, err error) {
	err = global.GVA_DB.Create(&pt).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processTypeEnter, err
}

// process type update
func ProcessTypeUpdate(pt *admin.ProcessType) (processType *admin.ProcessType, err error) {
	err = global.GVA_DB.Raw("UPDATE pm_process_type_c SET name = ?,  valid_id = ?, change_by = ?, change_time = ? WHERE id = ? ", pt.Name, pt.Valid, pt.ChangeBy, datatime.CurrentTimestamp(), pt.ID).Scan(&processType).Error
	if err != nil {
		fmt.Println("err update:", err)
		return
	}
	return processType, err
}

func ProcessGet(processID int) (processEnter admin.Process, err error) {
	err = global.GVA_DB.Table("dmc_pm_process").First(&processEnter, processID).Error
	if err != nil {
		return
	}
	return processEnter, err
}

// process add
func ProcessAdd(process *admin.Process) (processEnter admin.Process, err error) {
	err = global.GVA_DB.Table("dmc_pm_process").Create(&process).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processEnter, err
}

// process update
func ProcessUpdate(process *admin.Process) (processEnter admin.Process, err error) {
	err = global.GVA_DB.Raw(`UPDATE dmc_pm_process SET name = ?, description = ?, state_entity_id = ?,layout =?, config = ?, process_type =?, change_by = ?, change_time =? WHERE id = ?`, process.Name, process.Description, process.StateEntityID, process.Layout, process.Config, process.ProcessType, process.ChangeBy, process.ChangeTime, process.ID).Scan(&processEnter).Error
	if err != nil {
		fmt.Println("err add:", err)
		return
	}
	return processEnter, err
}
