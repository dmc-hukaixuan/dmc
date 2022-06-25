package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"
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

func (*Richtext) ValueSet(fieldID int, object string, objectID int, value string) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Richtext) ValueGet() {

}

func (*Richtext) SearchSQLGet() {

}

func (*Richtext) EditFieldRender() {

}

func (*Richtext) EditFieldValueGet() {

}

func (*Richtext) SearchFieldRender() {

}

func (*Richtext) StatsFieldParameterBuild() {

}
