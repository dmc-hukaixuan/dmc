package admin

import (
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/ticket"
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
	//user_id, _ := c.Get("userID")
	// edit dynamic field
	if sd.SubAction == "edit" {
		roleID, _ := sd.Data["roleID"].(string)
		roleIDint, _ := strconv.Atoi(roleID)
		fmt.Println(" SubAction roleID :", roleIDint)
		roleEdit(roleIDint, c)
	} else if sd.SubAction == "save" {
		var roleData model.TicketState
		mapstructure.Decode(sd.Data, &roleData)
		//roleSave(roleData, user_id.(int))
	} else {
		roleList := user.RoleOverview()
		response.SuccessWithDetailed(gin.H{
			"overviewList": &roleList,
		}, "获取成功", c)
	}
}

/*
   state edit
*/
func (t *TicketStateAPI) TicketStateEdit(c *gin.Context) {
	var stateData model.TicketState
	stateID, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(" roleEdit  roleID: ", stateID)
	if stateID > 0 {
		fmt.Println("  get role : ", stateID)
		stateData = service.StateGet(stateID)
	}

	stateTypeList := service.StateTypeList()
	state := map[string]interface{}{
		"name": stateData.Name,
		"state_type": map[string]interface{}{
			"Options": stateTypeList,
			"Default": stateData.StateTypeID,
		},
		"tag": map[string]interface{}{
			"Options": map[string]string{},
			"Default": "",
		},
		"comment": stateData.Comment,
	}

	response.SuccessWithDetailed(gin.H{
		"overviewList": &state,
	}, "获取成功", c)
}

/*
	ticket state add or update
*/
func (t *TicketStateAPI) TicketStateSave(c *gin.Context) {
	var stateData model.TicketState
	_ = c.ShouldBindJSON(&stateData)
	fmt.Println("sd ", stateData)
	user_id, _ := c.Get("userID")

	state_id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(" userID ", user_id.(int), stateData.ID, ", ValidID:", stateData.ValidID)
	if state_id > 0 {
		stateData.ID = state_id
		stateData.ChangeBy = user_id.(int)
		stateData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.StateUpdate(stateData)
	} else {
		stateData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		stateData.CreateBy = user_id.(int)
		stateData.ChangeBy = user_id.(int)
		stateData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.StateAdd(stateData)
	}
	response.Success(c)
}

func (t *TicketStateAPI) TicketStateOverview(c *gin.Context) {
	stateList := service.StateListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": &stateList,
	}, "获取成功", c)
}
