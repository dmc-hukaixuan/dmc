package admin

import (
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/user"
	"dmc/kernel/service/admin/user"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type UserAPI struct {
}

//  field lib base deal
func (t *UserAPI) UserBase(c *gin.Context) {
	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)
	user_id, _ := c.Get("userID")
	// edit dynamic field
	if sd.SubAction == "edit" {
		roleID, _ := sd.Data["roleID"].(string)
		roleIDint, _ := strconv.Atoi(roleID)
		fmt.Println(" SubAction roleID :", roleIDint)
		roleEdit(roleIDint, c)
	} else if sd.SubAction == "save" {
		var roleData model.Role
		mapstructure.Decode(sd.Data, &roleData)
		roleSave(roleData, user_id.(int))
	} else {
		roleList := user.RoleOverview()
		response.SuccessWithDetailed(gin.H{
			"overviewList": &roleList,
		}, "获取成功", c)
	}
}
