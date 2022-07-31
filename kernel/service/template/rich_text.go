package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket/dynamicField"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Richtext struct{}

func (*Richtext) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.FieldData) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.Name != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          fieldName,
			Default:       "",
			FieldType:     "richtext",
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

func (*Richtext) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Richtext) ValueGet(fieldID int, object string, objectID int64) interface{} {
	values := service.ValueGet(fieldID, objectID)
	return values[0].ValueText
}

func (*Richtext) SearchSQLGet() {

}

func (*Richtext) EditFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField, value interface{}) model.FieldData {
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

func (*Richtext) EditFieldValueGet() {

}

func (*Richtext) SearchFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig) model.FieldData {
	fieldData := model.FieldData{
		Name:        fieldName,
		Default:     "",
		FieldType:   "text",
		Label:       fieldLabel,
		Placeholder: "",
		Display:     1,
	}
	return fieldData
}

func (*Richtext) StatsFieldParameterBuild() {

}
