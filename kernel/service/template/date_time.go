package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket/dynamicField"
	// "encoding/json"
)

type DataTime struct{}

func (*DataTime) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.FieldData) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.Name != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          fieldObject.Name,
			Default:       "",
			FieldType:     "date",
			Label:         fieldLabel,
			Placeholder:   "",
			Display:       fieldObject.Display,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   DynamicFieldConfig.HintContent,
			HintType:      DynamicFieldConfig.HintType,
			Width:         4,
			RegexError:    "",
			Regex:         "",
		}
	}

	return fieldData
}

func (*DataTime) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*DataTime) ValueGet(fieldID int, object string, objectID int64) interface{} {
	values := service.ValueGet(fieldID, objectID)
	return values[0].ValueDate
}

func (*DataTime) SearchSQLGet() {

}

func (*DataTime) EditFieldRender() {

}

func (*DataTime) EditFieldValueGet() {

}

func (*DataTime) SearchFieldRender() {

}

func (*DataTime) StatsFieldParameterBuild() {

}
