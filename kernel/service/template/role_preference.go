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

func (*RolePreference) ValueSet(fieldID int, object string, objectID int64, value interface{}) {

}

func (*RolePreference) ValueGet(fieldID int, object string, objectID int64) interface{} {
	value := []interface{}{}
	return value
}

func (*RolePreference) SearchSQLGet() {

}

func (*RolePreference) EditFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField, value interface{}) model.FieldData {
	var fieldData model.FieldData

	fieldData = model.FieldData{
		Name:          "title",
		Default:       "",
		FieldType:     "cascader",
		Label:         FieldObject.FieldKey,
		Placeholder:   "",
		Display:       FieldObject.Display,
		Impacts:       []string{},
		DependsOn:     []string{},
		PromptCode:    2,
		PromptMessage: "",
		HintMessage:   "",
		HintType:      2,
		Width:         24,
		RegexError:    "",
		Regex:         "",
	}
	return fieldData
}

func (*RolePreference) EditFieldValueGet() {

}

func (*RolePreference) SearchFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig) model.FieldData {
	fieldData := model.FieldData{
		Name:        fieldName,
		Default:     "",
		FieldType:   "text",
		Label:       fieldLabel,
		Options:     DynamicFieldConfig.PossibleValues,
		Placeholder: "",
		Display:     1,
		Multiple:    1,
	}
	return fieldData
}

func (*RolePreference) StatsFieldParameterBuild() {

}
