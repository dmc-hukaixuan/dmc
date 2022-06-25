package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	user "dmc/kernel/service/admin/user"
	//"encoding/json"
)

type Role struct{}

func (*Role) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	roleList := user.RoleList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = roleList
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "role",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "Role",
			Placeholder:   "ticket service",
			Display:       2,
			Options:       roleList,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}

	return fieldData
}

func (*Role) ValueSet() {

}

func (*Role) ValueGet() {

}

func (*Role) SearchSQLGet() {

}

func (*Role) EditFieldRender() {

}

func (*Role) EditFieldValueGet() {

}

func (*Role) SearchFieldRender() {

}

func (*Role) StatsFieldParameterBuild() {

}
