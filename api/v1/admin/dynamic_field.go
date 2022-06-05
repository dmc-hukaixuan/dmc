package admin

import (
	model "dmc/kernel/model/admin"
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	"dmc/kernel/service/admin"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	yaml "gopkg.in/yaml.v2"
)

type DynamicFieldApi struct {
	// BaseController
}

//  field lib base deal
func (df *DynamicFieldApi) DynmicFieldbase(c *gin.Context) {
	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)

	// edit dynamic field
	if sd.SubAction == "edit" {
		field_id, _ := sd.Data["field_id"].(string)
		templateID1, _ := strconv.Atoi(field_id)
		fmt.Println(" templateID", field_id)
		DynamicFieldEdit(templateID1, c)
	} else if sd.SubAction == "save" {
		DynamicFieldUpdate(sd.Data, c)
	} else if sd.SubAction == "delete" {

	} else {
		df, _ := admin.DynamicFieldListGet()
		println("df ", df)
		response.SuccessWithDetailed(gin.H{
			"overViewList": df,
		}, "获取成功", c)
	}
}

// eidt dynmaicfield screen
func DynamicFieldEdit(fieldKey int, c *gin.Context) {
	var dyConfig model.DynamicField
	// get dynamic field config
	if fieldKey > 0 {
		dyConfig = admin.DynamicFieldGet(fieldKey, "")
	}

	// get field data
	fieldData := map[string]interface{}{
		"id":    dyConfig.ID,
		"name":  dyConfig.Name,
		"label": dyConfig.Label,
		"field_type": map[string]interface{}{
			"default": dyConfig.FieldType,
			"options": map[string]string{
				"radio":      "Radio",
				"dropdown":   "Drop down",
				"cascader":   "Cascader",
				"text":       "Text",
				"text_area":  "Text area",
				"date":       "Date",
				"date_time":  "Date time",
				"number":     "Number",
				"attachment": "Attachment",
				"checkbox":   "Checkbox",
				"rich_text":  "Rich text",
				"score":      "score",
			},
		},
		"object_type": map[string]interface{}{
			"default": dyConfig.ObjectType,
			"options": map[string]string{
				"Ticket":         "Ticket",
				"User":           "User",
				"Role":           "Role",
				"FAQ":            "FAQ",
				"ITSMConfigItem": "ITSMConfigItem",
			},
		},
		"config": dyConfig.Config,
	}
	response.SuccessWithDetailed(gin.H{
		"detail": fieldData,
	}, "获取成功", c)
}

// add or update dynamic field
func DynamicFieldUpdate(df map[string]interface{}, c *gin.Context) {
	var dfData model.DynamicField
	mapstructure.Decode(df, &dfData)
	user_id, _ := c.Get("userID")

	config, _ := yaml.Marshal(df["config"])
	if dfData.ID > 0 {
		dfData.ChangeBy = user_id.(int)
		dfData.ConfigT = string(config)
		dfData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		admin.DynamicFieldUpdate(dfData)
	} else {
		dfData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		dfData.CreateBy = user_id.(int)
		dfData.ChangeBy = user_id.(int)
		dfData.ConfigT = string(config)
		dfData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		admin.DynamicFieldAdd(dfData)
	}
	response.Success(c)
}

func DynamicFieldDelete() {

}
