package ticket

import (
	model "dmc/kernel/model/admin"
	"encoding/json"
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

func (*Role) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
	var perference_data model.FieldData
	roleList := user.RoleList(1)
	json.Unmarshal([]byte(fieldObject.FieldPreference), &perference_data)
	perference_data.Options = roleList
	// fieldData = model.FieldData{
	// 	Name:          "title",
	// 	Default:       "",
	// 	FieldType:     "text",
	// 	Label:         fieldObject.FieldKey,
	// 	Placeholder:   "",
	// 	Display:       fieldObject.Display,
	// 	Impacts:       []string{},
	// 	DependsOn:     []string{},
	// 	PromptCode:    2,
	// 	PromptMessage: "",
	// 	HintMessage:   "",
	// 	HintType:      2,
	// 	Width:         4,
	// 	RegexError:    "",
	// 	Regex:         "",
	// }
	return perference_data
}

func (*Role) EditFieldValueGet() {

}

func (*Role) SearchFieldRender() model.FieldData {
	roleList := user.RoleList(1)
	fieldData := model.FieldData{
		Name:        "Role",
		Default:     "",
		FieldType:   "dropdown",
		Label:       "role",
		Placeholder: "",
		Options:     roleList,
		Display:     1,
	}
	return fieldData
}

func (*Role) StatsFieldParameterBuild() {

}
