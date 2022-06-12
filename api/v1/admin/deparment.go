package admin

import (
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/ticket"
	user "dmc/kernel/service/admin/user"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DepartmentApi struct {
	// BaseController
}

func (d *DepartmentApi) DepartmentAdd(c *gin.Context) {
	var deparmentData model.Deparment
	_ = c.ShouldBindJSON(&deparmentData)
	user_id, _ := c.Get("userID")
	deparmentData.ChangeBy = user_id.(int)
	deparmentData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	deparmentData.CreateBy = user_id.(int)
	deparmentData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
	_, err := user.DepartmentAdd(deparmentData)
	if err == nil {
		response.SuccessWithMessage("新增成功", c)
	} else {
		response.FailWithDetailed(err, "操作失败", c)
	}
}

// service edit
func (d *DepartmentApi) DepartmentEdit(c *gin.Context) {
	department_id, _ := strconv.Atoi(c.Param("id"))
	DepartmentData := user.DepartmentGet(department_id)
	departmentList := user.DepartmentList(department_id)
	UserList := user.UserList(1)

	data := map[string]interface{}{
		"id":   DepartmentData.ID,
		"name": DepartmentData.Name,
		"departmentUser": map[string]interface{}{
			"default": DepartmentData.DeparmentUserList,
			"options": UserList,
		},
		"parent_department": map[string]interface{}{
			"default":    DepartmentData.ParentDepartmentID,
			"options":    departmentList,
			"field_type": "tree",
		},
		"Street":     DepartmentData.Street,
		"Zip":        DepartmentData.Zip,
		"City":       DepartmentData.City,
		"Country":    DepartmentData.Country,
		"DistrictID": DepartmentData.DistrictID,
		"Url":        DepartmentData.Url,
		"resplible": map[string]interface{}{
			"default": DepartmentData.DeparmentUserList,
			"options": UserList,
		},
		"validID": DepartmentData.ValidID,
		"comment": DepartmentData.Comment,
	}
	response.SuccessWithDetailed(gin.H{
		"typeData": &data,
	}, "获取成功", c)
}

func (d *DepartmentApi) DepartmentSave(c *gin.Context) {
	var deparmentData model.Deparment
	_ = c.ShouldBindJSON(&deparmentData)
	fmt.Println("deparmentData ", deparmentData)
	user_id, _ := c.Get("userID")
	deparmentData.ChangeBy = user_id.(int)
	deparmentData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
	user.DepartmentUpdate(deparmentData)
	response.SuccessWithMessage("提交成功", c)
}

func (d *DepartmentApi) DepartmentOverview(c *gin.Context) {
	list := user.DepartmentListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": &list,
	}, "获取成功", c)
}
