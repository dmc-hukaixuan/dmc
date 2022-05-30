package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
)

/*
	方法说明：获取日历中的 sla_calender

	入参：
		calender int

	返回
		WorkingCalender
*/
func WorkingTimeGet(working model.SLACalender) (working_time_id int, err error) {

	// ask database
	err = global.GVA_DB.Table("sla_working_time").Create(&working).Error
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		panic("err ")
	}

	//db.Close()
	return working.ID, err
}

func WorkingTimeAdd(calender int) {

}

func WorkingTimeUpdate(calender int) {

}

func workingTimeList() {

}
