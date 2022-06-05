package admin

import (
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/ticket"
	service "dmc/kernel/service/admin/ticket"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	yaml "gopkg.in/yaml.v2"
)

type WorkingTimeApi struct {
}

//  field lib base deal
func (wt *WorkingTimeApi) WorkingTimebase(c *gin.Context) {
	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)
	// edit dynamic field
	if sd.SubAction == "edit" {
		templateID, _ := sd.Data["id"].(string)
		templateID1, _ := strconv.Atoi(templateID)
		fmt.Println(" templateID", templateID)
		workingTimeEdit(templateID1, c)
	} else if sd.SubAction == "save" {
		workingTimeSave(sd.Data, c)
	} else {

	}
}

// working time hours edit
func workingTimeEdit(wt_id int, c *gin.Context) {
	var wt model.WorkingTimeCalender
	// get dynamic field config
	if wt_id > 0 {
		wt = service.WorkingTimeGet(wt_id)
		fmt.Println("wt_id :", wt_id)
	}
	var wc model.WorkingCalender
	fmt.Println("wt.WorkingTime :", wt.WorkingTime)
	yaml.Unmarshal([]byte(wt.WorkingTime), &wc)
	//
	fieldData := map[string]interface{}{
		"id":   wt.ID,
		"name": wt.Name,
		"timeZone": map[string]interface{}{
			"default": wt.TimeZone,
			"options": map[string]string{
				"UTC":           "UTC",
				"Asia/Shanghai": "Asia/Shanghai",
			},
		},
		"weekDayStart":    wt.WeekDayStart,
		"validID":         wt.ValidID,
		"workingTime":     wc.WorkingHours,
		"extraWorkingDay": wc.ExtraWorkingDay,
		"vacationDays":    wc.VacationDays,
		"comment":         wt.Comment,
	}
	response.SuccessWithDetailed(gin.H{
		"detail": fieldData,
	}, "获取成功", c)
}

func workingTimeSave(sd map[string]interface{}, c *gin.Context) {
	var wtc model.WorkingTimeCalender
	mapstructure.Decode(sd, &wtc)

	workingTime := map[string]interface{}{
		"extraWorkingDay": sd["extraWorkingDay"],
		"vacationDays":    sd["vacationDays"],
		"workingTime":     sd["workingTime"],
		"FormatOtherData": map[string]interface{}{
			"ExtraWorkingDay": "",
			"VacationDays":    "",
		},
	}
	config, _ := yaml.Marshal(workingTime)

	fmt.Println("wtc.ID ", wtc.ID)
	if wtc.ID > 0 {
		wtc.ChangeBy = 1
		wtc.WorkingTime = string(config)
		wtc.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.WorkingTimeUpdate(wtc)
	} else {
		wtc.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		wtc.CreateBy = 1
		wtc.ChangeBy = 1
		wtc.WorkingTime = string(config)
		wtc.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.WorkingTimeAdd(wtc)
	}
}

func (wt *WorkingTimeApi) WorkingTimeList(c *gin.Context) {
	wtl := service.WorkingTimeListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": wtl,
	}, "获取成功", c)
}
