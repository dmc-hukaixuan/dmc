package template

import (
	"dmc/global"
	model "dmc/kernel/model/admin"
)

type TicketTemplate struct{}

var TicketTemplateN = new(TicketTemplate)

/*
Get ticket template list

returns the ticket template list.
*/
func (p *TicketTemplate) TicketTemplateList() (list interface{}, total int64, err error) {
	var processType model.TemplateData
	// make a sql string
	selectSQL := `SELECT id, name, web, mobile, describe, valid_id, icon, color, type,
					display_type, create_time, create_by, change_time, change_by FROM  pm_process_type_c ptc`
	err = global.GVA_DB.Raw(selectSQL).Scan(&processType).Error
	if err != nil {
		return
	}
	return processType, total, err
}

/*
	Get ticket template detail data

	returns the ticket template data.
*/
func (p *TicketTemplate) TicketTemplateGet(templateid int) (list interface{}, total int64, err error) {
	var templateData model.TemplateData
	// make a sql string
	selectSQL := `SELECT id, name, web, mobile, describe, valid_id, icon, color, type,
				 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
	err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateData).Error
	if err != nil {
		return
	}
	// get ticket template link role
	role := templateRoleGet(templateid)
	// get temlate preference
	templateRoleGet(templateid)
	return templateData, total, err
}

/*
	Get ticket template detail data

	returns the ticket template data.
*/
func (p *TicketTemplate) TicketTemplateAdd(templateid int) (list interface{}, total int64, err error) {
	var templateData model.TemplateData
	// make a sql string
	selectSQL := `SELECT id, name, web, mobile, describe, valid_id, icon, color, type,
				 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
	err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateData).Error
	if err != nil {
		return
	}
	return templateData, total, err
}

/*
	Get ticket template detail data

	returns the ticket template data.
*/
func (p *TicketTemplate) TicketTemplateUpdate(templateid int) (list interface{}, total int64, err error) {
	var templateData model.TemplateData
	// make a sql string
	selectSQL := `SELECT id, name, web, mobile, describe, valid_id, icon, color, type,
				 display_type, create_time, create_by, change_time, change_by FROM dmc_ticket_template WHERE id = ? `
	err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateData).Error
	if err != nil {
		return
	}
	return templateData, total, err
}

/*
	delete ticket template

	returns the ticket template data.
*/
func (p *TicketTemplate) TicketTemplateDelete(templateid int) {
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
func templatePreferenceFieldAdd(templateid int) {
	var fieldInfo []model.TemplateField
	// do loop, push role data in template struct
	//

	roles := []string{}
	order := 1
	for _, v := range roles {
		fieldInfo.push(
			model.TemplateField{
				TemplateID:      templateid,
				RoleID:          v,
				FieldKey:        "",
				FieldLabel:      "",
				FieldOrder:      order,
				FieldObject:     "",
				FieldWidth:      "",
				FieldPreference: "",
			},
		)
		order++
	}
	// do bulk insert
	global.GVA_DB.Table("dmc_ticket_template_role").Create(&fieldInfo).Error
}

/*
	delete the template preference field
*/
func templatePreferenceFieldDelete(templateid int) (err error) {
	selectSQL := `DELETE FROM dmc_ticket_template_prefs WHERE template_id = ? `
	// ask database
	err = global.GVA_DB.Raw(selectSQL, templateid).Error
	if err != nil {
		return
	}
	return err
}

/*
	template field get
*/
func templatePreferenceFieldGet(templateid int) (err error) {
	var templateFields model.TemplateField
	// make a sql string
	selectSQL := `SELECT template_id, field_key, field_type, field_label, field_order, field_object, field_width 
				  field_preference, FROM dmc_ticket_template_prefs WHERE template_id = ? ORDER BY field_order`
	err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateFields).Error
	if err != nil {
		return
	}
	return templateFields, err
}

/*
	template role get
*/
func templateRoleGet(templateid int) (list interface{}, err error) {
	var templateRoles model.TemplateRole
	// make a sql string
	selectSQL := `SELECT template_id, role_id FROM dmc_ticket_template_role WHERE template_id = ? `
	err = global.GVA_DB.Raw(selectSQL, templateid).Scan(&templateRoles).Error
	if err != nil {
		return
	}
	return templateRoles, err
}

/*
	link template and role
*/
func templateRoleAdd(templateid int, roles []int) {
	var templateRoles []model.TemplateRole
	// do loop, push role data in template struct
	for _, v := range roles {
		templateRoles.push(
			model.TemplateRole{
				TemplateID: templateid,
				RoleID:     v,
			},
		)
	}
	// do bulk insert
	global.GVA_DB.Table("dmc_ticket_template_role").Create(&templateRoles).Error
}

/*
	delete ticket template link role data
*/
func templateRoleDelete(templateid int) (err error) {
	selectSQL := `DELETE FROM dmc_ticket_template_role WHERE template_id = ? `
	// ask database
	err = global.GVA_DB.Raw(selectSQL, templateid).Error
	if err != nil {
		return
	}
	return err
}
