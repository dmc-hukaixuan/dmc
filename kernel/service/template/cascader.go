package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Cascader struct{}

func (*Cascader) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if FieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(FieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
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
	}

	return fieldData
}

func (*Cascader) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueText: value,
	})

	service.ValueSet(fieldID, objectID, values)
}

func (*Cascader) ValueGet(fieldID string, object string, objectID int64) {

}

func (*Cascader) SearchSQLGet() {

}

func (*Cascader) EditFieldRender() {

}

func (*Cascader) EditFieldValueGet() {

}

func (*Cascader) SearchFieldRender() {

}

func (*Cascader) StatsFieldParameterBuild() {

}
