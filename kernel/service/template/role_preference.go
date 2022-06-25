package template

import (
	model "dmc/kernel/model/admin"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type RolePreference struct{}

func (*RolePreference) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.FieldData) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.Name != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          fieldName,
			Default:       "",
			FieldType:     "dropdown",
			Label:         fieldLabel,
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
	}

	return fieldData
}

func (*RolePreference) ValueSet() {

}

func (*RolePreference) ValueGet() {

}

func (*RolePreference) SearchSQLGet() {

}

func (*RolePreference) EditFieldRender() {

}

func (*RolePreference) EditFieldValueGet() {

}

func (*RolePreference) SearchFieldRender() {

}

func (*RolePreference) StatsFieldParameterBuild() {

}
