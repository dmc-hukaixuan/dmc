package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
	方法说明：获取日历中的 sla_calender

	入参：
		calender int

	返回
		WorkingCalender
*/
func WorkingTimeGet(wt_id int) (wt model.WorkingTimeCalender) {
	// ask database
	err := global.GVA_DB.Table("sla_working_time").Where("id = ?", wt_id).First(&wt).Error
	if err != nil {
		panic(err)
	}
	return wt
}

func WorkingTimeAdd(wtc model.WorkingTimeCalender) (working_time_id int, err error) {
	err = global.GVA_DB.Table("sla_working_time").Create(&wtc).Error
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		panic("err ")
	}

	//db.Close()
	return wtc.ID, err
}

func WorkingTimeUpdate(wtc model.WorkingTimeCalender) (working_time_id int, err error) {
	err = global.GVA_DB.Table("sla_working_time").Where("id = ?", wtc.ID).Model(&wtc).Omit("create_by", "create_time").Updates(wtc).Error
	if err != nil {
		return
	}
	return wtc.ID, err
}

func WorkingTimeListGet(validID int) (wtc []model.WorkingTimeCalender) {
	selectSQL := `SELECT wt.id as id, wt.name AS name, wt.time_zone AS time_zone, wt.week_day_start AS week_day_start,
					wt.valid_id AS valid_id, wt.comments AS comments, wt.time_zone AS  time_zone,
					wt.create_by AS create_by, wt.change_by AS change_by,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, wt.create_time AS create_time, wt.change_time AS change_time
					FROM sla_working_time wt LEFT JOIN users u ON u.id = wt.create_by LEFT JOIN users u1 ON u1.id = wt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&wtc).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return wtc
}

/*
	get all queues

    WorkingTimeList := $WorkingTimeList->WorkingTimeList();

    WorkingTimeList := $WorkingTimeList->WorkingTimeList( Valid => 1 );
*/
func WorkingTimeList(validID int) map[string]string {
	var sla []model.SLA
	selectSQL := `SELECT id, name FROM sla_working_time`
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	err := global.GVA_DB.Raw(selectSQL).Scan(&sla).Error
	workingTimeList := map[string]string{}
	if err != nil {
		return workingTimeList
	}
	// do loop, build a json string
	for _, v := range sla {
		workingTimeList[strconv.Itoa(v.ID)] = v.Name
	}
	return workingTimeList
}
