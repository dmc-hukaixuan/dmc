package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Checkbox struct{}

func (*Checkbox) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          fieldName,
			Default:       "",
			FieldType:     "checkbox",
			Label:         fieldLabel,
			Placeholder:   "",
			Display:       1,
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
	}

	return fieldData
}

func (*Checkbox) ValueSet(fieldID int, object string, objectID int, value string) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Checkbox) ValueGet() {

}

func (*Checkbox) SearchSQLGet() {

}

func (*Checkbox) EditFieldRender() {

}

func (*Checkbox) EditFieldValueGet() {

}

func (*Checkbox) SearchFieldRender() {

}

func (*Checkbox) StatsFieldParameterBuild() {

}
