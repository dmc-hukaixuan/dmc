package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	user "dmc/kernel/service/admin/user"
	//"encoding/json"
)

type OwnerBase struct{}

func (*OwnerBase) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	userList := user.UserList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = userList
		return fieldData
	} else {
		//differenct field
		label := ""
		NameLabel := ""
		display := 2
		if fieldType == "customer" {
			label = "customer user"
			NameLabel = "customeruser"
		} else if fieldType == "owner" {
			label = "Owner"
			NameLabel = "owner"
		} else {
			label = "Responsible"
			NameLabel = "responsible"
			display = 1
		}
		// field detail info
		fieldData = model.FieldData{
			Name:          NameLabel,
			Default:       "",
			FieldType:     "dropdown",
			Label:         label,
			Placeholder:   "",
			Display:       display,
			Options:       userList,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}
	return fieldData
}
func (*OwnerBase) ValueSet() {

}

func (*OwnerBase) ValueGet() {

}

func (*OwnerBase) SearchSQLGet() {

}

func (*OwnerBase) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
	var fieldData model.FieldData

	fieldData = model.FieldData{
		Name:          "title",
		Default:       "",
		FieldType:     "text",
		Label:         fieldObject.FieldKey,
		Placeholder:   "",
		Display:       fieldObject.Display,
		Impacts:       []string{},
		DependsOn:     []string{},
		PromptCode:    2,
		PromptMessage: "",
		HintMessage:   "",
		HintType:      2,
		Width:         4,
		RegexError:    "",
		Regex:         "",
	}
	return fieldData
}

func (*OwnerBase) EditFieldValueGet() {

}

func (*OwnerBase) SearchFieldRender() {

}

func (*OwnerBase) StatsFieldParameterBuild() {

}
