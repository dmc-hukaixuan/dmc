package admin

import (
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/ticket"
	service "dmc/kernel/service/admin/ticket"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SLAApi struct {
	// BaseController
}

/**
sla edit
*/
func (s *SLAApi) SLAEdit(c *gin.Context) {
	sla_id, _ := strconv.Atoi(c.Param("id"))
	SLAData := service.SLAGet(sla_id)
	serviceList := service.ServiceList(1)
	tagList := service.TagList(0)
	data := map[string]interface{}{
		"id":            SLAData.ID,
		"name":          SLAData.Name,
		"calendar_name": SLAData.CalendarName,
		"serviceList": map[string]interface{}{
			"default": SLAData.ServiceList,
			"options": serviceList,
		},
		"taglist": map[string]interface{}{
			"default": SLAData.TagList,
			"options": tagList,
		},
		"indicatorconfig": SLAData.IndicatorConfig,
		"internalNote":    SLAData.InternalNote,
		"externalNote":    SLAData.ExternalNote,
		"validID":         SLAData.ValidID,
		"comment":         SLAData.Comment,
	}
	response.SuccessWithDetailed(gin.H{
		"slaData": &data,
	}, "获取成功", c)
}

func (s *SLAApi) SLAUpdate(c *gin.Context) {
	var slaData model.SLA
	_ = c.ShouldBindJSON(&slaData)
	user_id, _ := c.Get("userID")
	slaData.ChangeBy = user_id.(int)
	slaData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
	service.SLAUpdate(slaData)
}

// sla add
func (s *SLAApi) SLAAdd(c *gin.Context) {
	var slaData model.SLA
	_ = c.ShouldBindJSON(&slaData)
	user_id, _ := c.Get("userID")
	slaData.ChangeBy = user_id.(int)
	slaData.CreateBy = user_id.(int)
	slaData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	slaData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
	service.SLAAdd(slaData)
}

func (s *SLAApi) SLAOverview(c *gin.Context) {
	tagList := service.SLAListGet(0)
	response.SuccessWithDetailed(gin.H{
		"overviewList": &tagList,
	}, "获取成功", c)
}
