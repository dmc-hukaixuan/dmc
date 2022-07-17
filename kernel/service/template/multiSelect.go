package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket/dynamicField"

	// model "dmc/kernel/model/ticket"
	// "encoding/json"
	"fmt"
)

type MultiSelect struct{}

func (*MultiSelect) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	fmt.Println("DynamicFieldConfig ++++++++++++++++++++++++++", DynamicFieldConfig.PossibleValues)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:                 fieldName,
			Default:              "",
			FieldType:            "MultiSelect",
			Label:                fieldLabel,
			Placeholder:          "",
			Display:              fieldObject.Display,
			Impacts:              []string{},
			DependsOn:            []string{},
			Options:              DynamicFieldConfig.PossibleValues,
			OptionsValueComments: DynamicFieldConfig.PossibleComments,
			PromptCode:           2,
			PromptMessage:        "",
			HintMessage:          "",
			HintType:             2,
			Width:                24,
		}
	}

	return fieldData
}

func (*MultiSelect) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*MultiSelect) ValueGet(fieldID int, object string, objectID int64) interface{} {
	values := service.ValueGet(fieldID, objectID)
	value := []interface{}{}
	// tempalte
	for _, v := range values {
		value = append(value, v.ValueDate)
	}
	return value
}

func (*MultiSelect) SearchSQLGet() {

}

func (*MultiSelect) EditFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField, value interface{}) model.FieldData {
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

func (*MultiSelect) EditFieldValueGet() {

}

func (*MultiSelect) SearchFieldRender() {

}

func (*MultiSelect) StatsFieldParameterBuild() {

}
