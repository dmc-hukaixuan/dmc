package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"

	// model "dmc/kernel/model/ticket"
	// "encoding/json"
	"fmt"
)

type Dropdown struct{}

func (*Dropdown) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
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
			FieldType:            "dropdown",
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

func (*Dropdown) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueDate: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Dropdown) ValueGet() {

}

func (*Dropdown) SearchSQLGet() {

}

func (*Dropdown) EditFieldRender() {

}

func (*Dropdown) EditFieldValueGet() {

}

func (*Dropdown) SearchFieldRender() {

}

func (*Dropdown) StatsFieldParameterBuild() {

}
