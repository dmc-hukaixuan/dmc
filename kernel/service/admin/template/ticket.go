package template

import (
    "dmc/global"
    model "dmc/kernel/model/admin"
    "encoding/json"
    "fmt"
)

type TicketTemplate struct{}

var TicketTemplateN = new(TicketTemplate)

/*
Get ticket template list

returns the ticket template list.
*/
func TicketTemplateList() (templateList []model.TemplateData, err error) {
    // make a sql string
    selectSQL := `SELECT tt.id, tt.name, tt.web, tt.mobile, tt.describes, tt.valid_id, tt.icon, tt.color, tt.type,
                    tt.display_type, tt.create_time, u1.full_name AS create_by_name, tt.change_time, u2.full_name AS change_by FROM dmc_ticket_template tt
                    left join users u1 on tt.create_by = u1.id
                    left join users u2 on tt.change_by = u2.id`
    err = global.GVA_DB.Raw(selectSQL).Scan(&templateList).Error
    if err != nil {
        return
    }
    return templateList, err
}

/*
   Get ticket template detail data

   returns the ticket template data.
*/
func TicketTemplateGet(templateid int) (list model.TemplateData, tlf map[string]model.TemplateField, err error) {
    var templateData model.TemplateData
    // make a sql string
    selectSQL := `SELECT id, name, web, mobile, describes, valid_id, icon, color, type,
                 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
    err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateData).Error
    if err != nil {
        return
    }
    // get ticket template link role
    role, _ := templateRoleGet(templateid)
    // get temlate preference
    preference, _ := templatePreferenceFieldGet(templateid)
    fieldOrder := []string{}
    fmt.Println("preference---------------- ", preference)
    templateField := map[string]model.TemplateField{}
    for _, v := range preference {

        // json.Unmarshal([]byte(v.FieldPreference.(string)), v.FieldPreference)
        fieldOrder = append(fieldOrder, v.FieldKey)
        templateField[v.FieldKey] = v
    }
    templateData.Roles = role
    templateData.FieldOrder = fieldOrder

    fmt.Println("templateData---------------- ", templateField)
    return templateData, templateField, err
}

/*
   Get ticket template detail data

   returns the ticket template data.
*/
func TicketTemplateAdd(td model.TemplateData) (templateID int, err error) {

    // make a sql string
    err = global.GVA_DB.Table("dmc_ticket_template").Create(&td).Error
    fmt.Println(" err ---", err)
    if err != nil {
        return
    }
    fmt.Println("td.IDv:    ===== ", td.ID)
    // add ticket field data to db
    templatePreferenceFieldAdd(td.ID, td)
    // add template link role to db
    templateRoleAdd(td.ID, td.Roles)
    return td.ID, err
}

/*
   Get ticket template detail data

   returns the ticket template data.
*/
func TicketTemplateUpdate(td model.TemplateData) (list interface{}, err error) {
    var tdata model.TemplateData
    // make a sql string
    err = global.GVA_DB.Table("dmc_ticket_template").Where("id = ?", td.ID).Model(&tdata).Omit("create_by", "create_time").Updates(td).Error
    if err != nil {
        return
    }
    fmt.Println("err  -------------------", err)
    // first remove template lin preference field and remove role
    templatePreferenceFieldDelete(td.ID)
    templateRoleDelete(td.ID)
    // add ticket field data to db
    templatePreferenceFieldAdd(td.ID, td)
    // add template link role to db
    templateRoleAdd(td.ID, td.Roles)
    return td.ID, err
}

/*
   delete ticket template

   returns the ticket template data.
*/
func TicketTemplateDelete(templateid int) {
    // delete template link role and relationship
    templatePreferenceFieldDelete(templateid)
    templateRoleDelete(templateid)

    // make a sql string
    selectSQL := `DELETE FROM dmc_ticket_template WHERE id = ? `
    // ask database
    err := global.GVA_DB.Raw(selectSQL, templateid).Error
    if err != nil {
        return
    }
    return
}

/*
   ticket preferenct field add
*/
func templatePreferenceFieldAdd(templateID int, td model.TemplateData) (err error) {
    var fieldInfo []model.TemplateField
    fmt.Println(" td model.TemplateData   ", td.FieldOrder)
    // do loop, push role data in template struct
    order := 1
    for _, v := range td.FieldOrder {
        FieldPreference, _ := json.Marshal(td.FieldData[v])
        fieldInfo = append(fieldInfo,
            model.TemplateField{
                TemplateID:      td.ID,
                FieldKey:        v,
                FieldLabel:      td.FieldData[v].Label,
                FieldOrder:      order,
                FieldType:       td.FieldData[v].FieldType,
                FieldObject:     td.FieldData[v].FieldType,
                FieldWidth:      td.FieldData[v].Width,
                FieldPreference: string(FieldPreference),
            },
        )
        order++
    }
    // do bulk insert
    err = global.GVA_DB.Table("dmc_ticket_template_prefs").Create(&fieldInfo).Error
    return err
}

/*
   delete the template preference field
*/
func templatePreferenceFieldDelete(templateid int) (err error) {
    fmt.Println(" templateid ", templateid)
    selectSQL := `DELETE FROM dmc_ticket_template_prefs WHERE template_id = ? `
    // ask database
    err = global.GVA_DB.Exec(selectSQL, templateid).Unscoped().Error

    if err != nil {
        return
    }
    return err
}

/*
   template field get
*/
func templatePreferenceFieldGet(templateid int) (prefer []model.TemplateField, err error) {
    var templateFields []model.TemplateField
    // make a sql string
    selectSQL := `SELECT template_id, field_key, field_type, field_label, field_order, field_object, field_width,
                  field_preference FROM dmc_ticket_template_prefs WHERE template_id = ? ORDER BY field_order`
    err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateFields).Error
    if err != nil {
        return
    }
    return templateFields, err
}

/*
   template role get
*/
func templateRoleGet(templateid int) (roels []int, err error) {
    var templateRoles []model.TemplateRole
    // make a sql string
    selectSQL := `SELECT template_id, role_id FROM dmc_ticket_template_role WHERE template_id = ? `
    err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateRoles).Error
    if err != nil {
        return
    }
    roles := []int{}
    for _, v := range templateRoles {
        roles = append(roles, v.TemplateID)
    }
    return roles, err
}

/*
   link template and role
*/
func templateRoleAdd(templateid int, roles []int) (err error) {
    var templateRoles []model.TemplateRole
    if len(roles) == 0 {
        roles = []int{0}
    }
    // do loop, push role data in template struct
    for _, v := range roles {
        templateRoles = append(templateRoles,
            model.TemplateRole{
                TemplateID: templateid,
                RoleID:     v,
            },
        )
    }

    // do bulk insert
    err = global.GVA_DB.Table("dmc_ticket_template_role").Create(&templateRoles).Error
    return err
}

/*
   delete ticket template link role data
*/
func templateRoleDelete(templateid int) (err error) {
    selectSQL := `DELETE FROM dmc_ticket_template_role WHERE template_id = ? `
    // ask database
    err = global.GVA_DB.Exec(selectSQL, templateid).Error
    if err != nil {
        return
    }
    return err
}
