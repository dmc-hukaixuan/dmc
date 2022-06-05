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

type TicketSourceAPI struct {
}

//  field lib base deal
func (t *TicketTypeAPI) TicketSourcebase(c *gin.Context) {
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
func (t *TicketSourceAPI) TypeEdit(c *gin.Context) {
	type_id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("type_id ", type_id)
	var ticketsourceData model.TicketSource
	fmt.Println(" roleEdit  roleID: ", type_id)
	if type_id > 0 {
		fmt.Println("  get role : ", type_id)
		ticketsourceData = service.SourceGet(type_id)
	}
	typeData := map[string]interface{}{
		"id":      ticketsourceData.ID,
		"name":    ticketsourceData.Name,
		"tnstart": ticketsourceData.Icon,
		"validID": ticketsourceData.ValidID,
		"color":   ticketsourceData.Color,
		"icon":    ticketsourceData.Icon,
		"comment": ticketsourceData.Comment,
	}

	response.SuccessWithDetailed(gin.H{
		"typeData": &typeData,
	}, "获取成功", c)
}

/*
	ticket type add or update
*/
func TicketSourceSave(typeData model.TicketType, userID int) {
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
func (t *TicketSourceAPI) TicketSourceSave(c *gin.Context) {
	type_id, _ := strconv.Atoi(c.Param("id"))
	var ticketSource model.TicketSource
	_ = c.ShouldBindJSON(&ticketSource)
	user_id, _ := c.Get("userID")

	//mapstructure.Decode(sd.Data, &typeData)

	if type_id > 0 {
		ticketSource.ChangeBy = user_id.(int)
		ticketSource.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.SourceUpdate(ticketSource)
	} else {
		ticketSource.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		ticketSource.CreateBy = user_id.(int)
		ticketSource.ChangeBy = user_id.(int)
		ticketSource.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
		service.SourceAdd(ticketSource)
	}
	response.Success(c)
}

func (t *TicketSourceAPI) TicketSourceList(c *gin.Context) {
	roleList := service.SourceListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": &roleList,
	}, "获取成功", c)
}
