package admin

import (
    "dmc/kernel/model/common/response"
    model "dmc/kernel/model/ticket"
    service "dmc/kernel/service/admin/ticket"
    "fmt"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

type ServiceApi struct {
    // BaseController
}

// service add api
func (s *ServiceApi) ServiceAdd(c *gin.Context) {
    var serviceData model.Service
    _ = c.ShouldBindJSON(&serviceData)
    user_id, _ := c.Get("userID")
    serviceData.ChangeBy = user_id.(int)
    serviceData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    serviceData.CreateBy = user_id.(int)
    serviceData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
    service.ServiceAdd(serviceData)
}

// service edit
func (s *ServiceApi) ServiceEdit(c *gin.Context) {
    service_id, _ := strconv.Atoi(c.Param("id"))
    ServiceData := service.ServiceGet(service_id)
    serviceList := service.ServiceList(1)
    slaList := service.SLAList(1)
    data := map[string]interface{}{
        "id":   ServiceData.ID,
        "name": ServiceData.Name,
        "parentNode": map[string]interface{}{
            "default": service_id,
            "options": serviceList,
        },
        "sla": map[string]interface{}{
            "default": ServiceData.SLAList,
            "options": slaList,
        },
        "taglist": map[string]interface{}{
            "default": ServiceData.TagList,
            "options": slaList,
        },
        "internalNote": ServiceData.InternalNote,
        "externalNote": ServiceData.ExternalNote,
        "validID":      ServiceData.ValidID,
        "comment":      ServiceData.Comment,
    }
    response.SuccessWithDetailed(gin.H{
        "typeData": &data,
    }, "获取成功", c)
}

func (s *ServiceApi) ServiceUpdate(c *gin.Context) {
    var serviceData model.Service
    _ = c.ShouldBindJSON(&serviceData)
    user_id, _ := c.Get("userID")
    serviceData.ChangeBy = user_id.(int)
    serviceData.ChangeTime = time.Now().Format("2006-01-02 15:04:05")
    service.ServiceUpdate(serviceData)
}

func (s *ServiceApi) ServiceOverview(c *gin.Context) {
    fmt.Println(" servicelisdt ")
    list := service.ServiceListGet(0)
    response.SuccessWithDetailed(gin.H{
        "overviewList": &list,
    }, "获取成功", c)
}

func (s *ServiceApi) ServiceLinkTag(c *gin.Context) {

}
