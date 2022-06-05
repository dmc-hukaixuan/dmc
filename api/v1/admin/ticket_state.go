package admin

import (
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/user"
	service "dmc/kernel/service/admin/ticket"
	"dmc/kernel/service/admin/user"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type TicketStateAPI struct {
}

//  field lib base deal
func (t *TicketStateAPI) TicketStatebase(c *gin.Context) {
	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)
	fmt.Println("sd ", sd)
	user_id := 1

	// edit dynamic field
	if sd.SubAction == "edit" {
		roleID, _ := sd.Data["roleID"].(string)
		roleIDint, _ := strconv.Atoi(roleID)
		fmt.Println(" SubAction roleID :", roleIDint)
		roleEdit(roleIDint, c)
	} else if sd.SubAction == "save" {
		var roleData model.Role
		mapstructure.Decode(sd.Data, &roleData)
		roleSave(roleData, user_id)
	} else {
		roleList := user.RoleOverview()
		response.SuccessWithDetailed(gin.H{
			"overviewList": &roleList,
		}, "获取成功", c)
	}
}

/*
 role edit
*/
func TicketStateEdit(roleID int, c *gin.Context) {
	var roleData model.Role
	fmt.Println(" roleEdit  roleID: ", roleID)
	if roleID > 0 {
		fmt.Println("  get role : ", roleID)
		roleData = user.RoleGet(roleID)
	}

	calenderList := service.WorkingTimeList(1)
	userList := user.UserList(1)
	roleList := user.RoleList(1)
	roles := map[string]interface{}{
		"name": roleData.Name,
		"parentRole": map[string]interface{}{
			"Options": roleList,
			"Default": "",
		},
		"defaultOwner": map[string]interface{}{
			"Options": userList,
			"Default": roleData.DefaultOwner,
		},
		"defaultResponsible": map[string]interface{}{
			"Options": userList,
			"Default": roleData.DefaultResponsible,
		},
		"calendar": map[string]interface{}{
			"Options": calenderList,
			"Default": roleData.CalendarName,
		},
		"systemAddress": map[string]interface{}{
			"Options": map[string]string{},
			"Default": "",
		},
		"tag": map[string]interface{}{
			"Options": map[string]string{},
			"Default": "",
		},
		"comment": roleData.Comments,
	}

	response.SuccessWithDetailed(gin.H{
		"overviewList": &roles,
	}, "获取成功", c)
}

/*
	ticket state add or update
*/
func TicketStateSave(roleData model.Role, userID int) {
	fmt.Println(" userID ", userID, roleData.ID, ", ValidID:", roleData.ValidID)
	if roleData.ID > 0 {
		roleData.ChangeBy = userID
		roleData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		user.RoleUpdate(roleData)
	} else {
		roleData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		roleData.CreateBy = userID
		roleData.ChangeBy = userID
		roleData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		user.RoleAdd(roleData)
	}
}
