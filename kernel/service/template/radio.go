package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Radio struct{}

func (*Radio) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	//test := *DynamicFieldConfig

	// get template orN
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:        fieldName,
			Default:     DynamicFieldConfig.DefaultValue,
			FieldType:   "radio",
			Label:       fieldLabel,
			Placeholder: "",

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

func (*Radio) ValueSet(fieldID int, object string, objectID int, value string) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Radio) ValueGet() {

}

func (*Radio) SearchSQLGet() {

}

func (*Radio) EditFieldRender() {

}

func (*Radio) EditFieldValueGet() {

}

func (*Radio) SearchFieldRender() {

}

func (*Radio) StatsFieldParameterBuild() {

}
