package template

import (
	model "dmc/kernel/model/admin"
	model_dynamicField "dmc/kernel/model/dynamicField"
	service "dmc/kernel/service/ticket"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Tree struct{}

func (*Tree) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:                 fieldName,
			Default:              "",
			FieldType:            "tree",
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
			Width:                4,
		}
	}

	return fieldData
}

func (*Tree) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
	values := []model_dynamicField.DynamicFieldValue{}
	values = append(values, model_dynamicField.DynamicFieldValue{
		FieldID:   fieldID,
		ObjectID:  objectID,
		ValueText: value,
	})
	service.ValueSet(fieldID, objectID, values)
}

func (*Tree) ValueGet() {

}

func (*Tree) SearchSQLGet() {

}

func (*Tree) EditFieldRender() {

}

func (*Tree) EditFieldValueGet() {

}

func (*Tree) SearchFieldRender() {

}

func (*Tree) StatsFieldParameterBuild() {

}
