package admin

import (
    model "dmc/kernel/model/admin"
    "dmc/kernel/model/common/request"
    "dmc/kernel/model/common/response"
    "dmc/kernel/service/admin"
    "time"

    templateObject "dmc/kernel/service/admin/template"
    dynamicFieldObject "dmc/kernel/service/template"
    ticketField "dmc/kernel/service/template/ticket"
    "fmt"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/mitchellh/mapstructure"
)

type TicketTemplateApi struct {
    // BaseController
}

// ticket template base
func (p *TicketTemplateApi) Base(c *gin.Context) {
    var sd request.SubActionData
    user_id, _ := c.Get("userID")
    _ = c.ShouldBindJSON(&sd)
    fmt.Println("sd -------1111---:", sd)
    if sd.SubAction == "edit" {
        templateID, _ := sd.Data["templateID"].(string)
        templateID1, _ := strconv.Atoi(templateID)
        fmt.Println(" templateID", templateID, user_id.(int))
        TicketTemplateEdit(templateID1, c)
    } else if sd.SubAction == "save" {
        var td model.TemplateData
        mapstructure.Decode(sd.Data, &td)
        ticketTemplateSave(td, c)
    } else if sd.SubAction == "delete" {
        // delete tempalte
        //ticketTemplateDelete(1)
    } else {
        templateList, _ := templateObject.TicketTemplateList()
        response.SuccessWithDetailed(gin.H{
            "templateList": &templateList,
        }, "获取成功", c)
    }
}

// get ticket template data ,for add or edit a template
func TicketTemplateEdit(templateID int, c *gin.Context) {
    // get ticket template data
    templateData, templateFieldList, _ := templateObject.TicketTemplateGet(templateID)
    // template base field

    // ticket base field
    baseField := map[string]interface{}{
        "name":        templateData.Name,
        "description": templateData.DisplayType,
        "valid":       templateData.ValidID,
        "showLocation": &model.FieldData{
            Default: "",
            Options: map[string]string{
                "mobile": "Show Mobile",
                "web":    "Show Web",
            },
        },
        "displayType": &model.FieldData{
            Default: templateData.DisplayType,
            Options: map[string]string{
                "1": "show form, edit field",
                "2": "not show form",
                "3": "open dailog",
            },
        },
        "templateType": &model.FieldData{
            Default: templateData.Type,
            Options: map[string]string{
                "create": "Create Normal Ticket",
                "const":  "Ticket Detail Show",
                "deal":   "Processing Ticket",
            },
        },
        "roles": &model.FieldData{
            Default:   templateData.Roles,
            FieldType: "dropdown",
            Options: map[string]string{
                "1": "Valid",
                "2": "Invalid",
            },
        },
        "color": templateData.Color,
        "icon":  templateData.Icon,
    }
    //fmt.Println(" baseField", baseField)
    // dynamic field
    templateFieldData, fieldOrder := allticketField(templateID, templateFieldList, templateData)

    response.SuccessWithDetailed(gin.H{
        "base": baseField,
        //	"filter":         allticketField(),
        "templateField":      templateFieldData,
        "templateFieldOrder": fieldOrder,
        "ticketRequired":     "",
    }, "获取成功", c)

}

// get ticket template data ,for add or edit a template
func (p *TicketTemplateApi) TicketTemplateGet(c *gin.Context) {
    // get ticket template

    // base field

    // dynamic field
    // DynamicFieldList("Ticket")
    response.SuccessWithDetailed(gin.H{
        "baseField":           [...]string{"name", "description", "processState", "processType"},
        "filterField":         "",
        "templateField":       "",
        "ticketRequiredField": "",
        "":                    "",
    }, "获取成功", c)
}

// add copy etc...
func (p *TicketTemplateApi) TicketTemplateUpdate(c *gin.Context) {
    // parse template data
}

// save ticket template
func ticketTemplateSave(td model.TemplateData, c *gin.Context) {
    // td
    if td.ID != 0 {
        fmt.Println("td -6666-- ", td.Name)
        td.Changetime = time.Now().Format("2006-01-02 15:04:05")
        td.Changeby = 1
        if _, err := templateObject.TicketTemplateUpdate(td); err != nil {
            fmt.Println("ProcessUpdate  err :", err)
            response.FailWithMessage("更新失败", c)
        } else {
            // fmt.Println("list ----------------:", typelist, "total:", typelist)
            response.SuccessWithMessage("更新成功", c)
        }
    } else {
        // add template to db
        td.Createtime = time.Now().Format("2006-01-02 15:04:05")
        td.Createby = 1
        td.Changetime = time.Now().Format("2006-01-02 15:04:05")
        td.Changeby = 1
        templateID, _ := templateObject.TicketTemplateAdd(td)
        fmt.Println("templateID: ", templateID)
    }
}

// ticket template, add copy etc...
func ticketTemplateUpdate() {

}

// remove ticket template
func ticketTemplateDelete(templateID int) {

}

// get ticket template list data
// func ticketTemplateList() {
// 	templateList, _ := templateObject.TicketTemplateList()
// 	fmt.Println("templateList ", templateList)
// }

// get all field
func allticketField(templateID int, templateFieldList map[string]model.TemplateField, templatedata model.TemplateData) (map[string]model.FieldData, []string) {
    fieldInfo := map[string]model.FieldData{}
    templateField := []string{"title", "type", "owner", "customeruser", "priority", "role", "Body", "state", "source", "service", "sla"}
    // ticketFieldObject := ticketField.TicketStartandField()
    if templateID > 0 {
        templateField = templatedata.FieldOrder
        fieldInfo = templatedata.FieldData
    }

    // get ticket each field detail info
    for _, v := range templateField {
        fieldObject := ticketField.TicketStartandField(v).TemplateEditRender(v, templateFieldList[v])
        fieldInfo[v] = fieldObject
    }
    // get dynamic field list
    df, _ := admin.DynamicFieldList("Ticket")
    //fmt.Println("df list :", df)
    // do loop dynamicfield list
    for _, v := range df {
        //fmt.Println("v0", v)
        field := dynamicFieldObject.DynamicField(v.FieldType).TemplateEditRender(v.Name, v.Label, &v.Config, templateFieldList[v.Name])
        fieldInfo["DynamicField_"+v.Name] = field
    }

    return fieldInfo, templateField
}
