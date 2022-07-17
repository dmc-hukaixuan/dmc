package ticket

import (
    modelAdmin "dmc/kernel/model/admin"
    "dmc/kernel/model/common/response"
    model "dmc/kernel/model/ticket"
    dynamicField "dmc/kernel/service/admin"
    processMB "dmc/kernel/service/admin/process"
    templateObject "dmc/kernel/service/admin/template"
    userService "dmc/kernel/service/admin/user"
    dynamicFieldObject "dmc/kernel/service/template"
    ticketField "dmc/kernel/service/template/ticket"
    service "dmc/kernel/service/ticket"
    "dmc/kernel/service/ticket/article"
    "dmc/kernel/service/ticket/process"
    "dmc/kernel/util/uploadCache"
    "fmt"
    "regexp"
    "strconv"

    "github.com/gin-gonic/gin"
)

type TemplateAPI struct{}

func (*TemplateAPI) TicketTemplateList(c *gin.Context) {
    // get user id
    user_id, _ := c.Get("userID")
    // get user roles
    // get user role
    userTemplate := userService.UserCreateTemplateList(user_id.(int))
    // get roles permission template
    // get template link proces node
    template := processMB.TemplateNodeList()

    processList, _, _ := processMB.ProcessList()
    createTemplateList := map[string]map[string]interface{}{}
    for _, v := range processList {
        if v.StateEntityID == 1 {
            var td []model.TemplateData
            for _, tl := range template {
                //
                if tl.ProcessID == v.EntityID {
                    if value, ok := userTemplate[tl.TemplateID]; ok {
                        td = append(td, value)
                    }
                }
            }
            createTemplateList[v.ProcessType] = map[string]interface{}{
                "processName":     v.Name,
                "processEntityID": v.EntityID,
                "processID":       v.ID,
                "firstNodeID":     v.FirstNodeId,
                "templateList":    td,
            }
        }
    }
    // get process
}

func (*TemplateAPI) TicketTemplateGet(c *gin.Context) {
    // get formid
    FormID := uploadCache.WebUploadCache().FormIDCreate()
    templateID, _ := strconv.Atoi(c.Param("id"))

    // get ticket template data
    templateData, templateFieldList, _ := templateObject.TicketTemplateGet(templateID)
    fmt.Println("templateData ", templateData)
    reg := regexp.MustCompile(`^dynamicField_`)
    // get all dynamicField id
    DynamicField := dynamicField.DynamicFieldNameList("Ticket")

    fieldInfo := map[string]modelAdmin.FieldData{}
    // get ticket each field detail info
    for _, field := range templateData.FieldOrder {
        if reg.MatchString(field) {
            v, _ := DynamicField[field]
            fieldObject := dynamicFieldObject.DynamicField(v.FieldType).EditFieldRender(v.Name, v.Label, &v.Config, templateFieldList[v.Name], "")
            fieldInfo[field] = fieldObject
        } else {
            fieldObject := ticketField.TicketStartandField(field).EditFieldRender(field, templateFieldList[field], "")
            fieldInfo[field] = fieldObject
        }
    }

    // return ticket template
    response.SuccessWithDetailed(gin.H{
        "templateField": fieldInfo,
        "formID":        FormID,
        "fieldOrder":    templateData.FieldOrder,
    }, "获取成功", c)
}

func TicketCreate(c *gin.Context) {
    var ticketBaseData model.TicketBaseData
    _ = c.ShouldBindJSON(&ticketBaseData)

    user_id, _ := c.Get("userID")

    // create new ticket, do db insert
    ticket_id, _ := service.TicketCreate(ticketBaseData)

    ticketData := make(map[string]interface{}) //注意该结构接受k的内容
    c.BindJSON(&ticketData)

    // get all dynamicField id
    DynamicField := dynamicField.DynamicFieldNameList("Ticket")
    reg := regexp.MustCompile(`^dynamicField_`)
    // set ticket dynamic fields
    // cycle through the activated Dynamic Fields for this screen
    for k, v := range ticketData {
        if reg.MatchString(k) {
            //  set the value
            dynamicFieldObject.DynamicField(DynamicField[k].FieldType).ValueSet(DynamicField[k].ID, "Ticket", ticket_id, v)
        }
    }

    form := userService.UserGet(ticketData["customer"].(int))
    to := userService.RoleGet(ticketData["queue"].(int))

    SenderType := "agent"

    // article create
    // get pre loaded attachment
    articleID := article.ArticleCreate(
        model.ArticleDataMimeCreate{
            TicketID:             ticket_id,
            SenderType:           SenderType,
            IsVisibleForCustomer: ticketData["isVisibleForCustomer"].(int),
            From:                 form.FullName + " <" + form.Email + " >",
            To:                   to.Name,
            Subject:              ticketData["subject"].(string),
            Body:                 ticketData["body"].(string),
            MimeType:             "text/plain",
            Charset:              "charset=utf-8",
            UserID:               user_id.(int),
        },
    )
    // write attachments and form id
    article.ArticleWriteAttachment(ticketData["formID"].(string), articleID, user_id.(int), ticket_id)

    // remove pre submited attachments
    uploadCache.WebUploadCache().FormIDRemoveFile(ticketData["formID"].(string))

    process.ProcessTransition(user_id.(int), ticket_id)

    // trgger event

    // link tickets
    // SplitLinkType， LinkType， Direction
    if splitTicketID, ok := ticketData["splitTicketID"]; ok {
        fmt.Println("splitTicketID :", splitTicketID)
    }
    // check link ticket permission

}
