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
)

type TicketTypeAPI struct {
}

//  field lib base deal
func (t *TicketTypeAPI) TicketTypebase(c *gin.Context) {
	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)
	user_id, _ := c.Get("userID")
	// edit dynamic field
	if sd.SubAction == "edit" {
		// type_id, _ := sd.Data["roleID"].(string)
		// typeID, _ := strconv.Atoi(type_id)
		// TypeEdit(typeID, c)
	} else if sd.SubAction == "save" {
		var typeData model.TicketType
		mapstructure.Decode(sd.Data, &typeData)
		TicketTypeSave(typeData, user_id.(int))
	} else {

	}
}

/*
 type edit
*/
func (t *TicketTypeAPI) TypeEdit(c *gin.Context) {
	type_id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("type_id ", type_id)
	var ticketTypeData model.TicketType
	fmt.Println(" roleEdit  roleID: ", type_id)
	if type_id > 0 {
		fmt.Println("  get role : ", type_id)
		ticketTypeData = service.TypeGet(type_id)
	}
	typeData := map[string]interface{}{
		"id":      ticketTypeData.ID,
		"name":    ticketTypeData.Name,
		"tnstart": ticketTypeData.Icon,
		"validID": ticketTypeData.ValidID,
		"color":   ticketTypeData.Color,
		"icon":    ticketTypeData.Icon,
		"comment": ticketTypeData.Comment,
	}

	response.SuccessWithDetailed(gin.H{
		"typeData": &typeData,
	}, "获取成功", c)
}

/*
	ticket type add or update
*/
func TicketTypeSave(typeData model.TicketType, userID int) {
	if typeData.ID > 0 {
		typeData.ChangeBy = userID
		typeData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.TypeUpdate(typeData)
	} else {
		typeData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		typeData.CreateBy = userID
		typeData.ChangeBy = userID
		typeData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.TypeAdd(typeData)
	}
}

/*
	ticket type add or update
*/
func (t *TicketTypeAPI) TypeSave(c *gin.Context) {
	type_id, _ := strconv.Atoi(c.Param("id"))
	var typeData model.TicketType
	_ = c.ShouldBindJSON(&typeData)
	user_id, _ := c.Get("userID")

	//mapstructure.Decode(sd.Data, &typeData)

	if type_id > 0 {
		typeData.ChangeBy = user_id.(int)
		typeData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.TypeUpdate(typeData)
	} else {
		typeData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		typeData.CreateBy = user_id.(int)
		typeData.ChangeBy = user_id.(int)
		typeData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.TypeAdd(typeData)
	}
	response.Success(c)
}

func (t *TicketTypeAPI) TypeList(c *gin.Context) {
	roleList := service.TypeListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": &roleList,
	}, "获取成功", c)
}
