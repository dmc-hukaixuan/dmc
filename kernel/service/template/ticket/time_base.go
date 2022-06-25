package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	//"encoding/json"
)

type TimeBase struct{}

func (*TimeBase) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		// incident start time
		// incident end time
		fieldData = model.FieldData{
			Name:          "title",
			Default:       "",
			FieldType:     "text",
			Label:         fieldObject.FieldKey,
			Placeholder:   "",
			Display:       fieldObject.Display,
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

func (*TimeBase) ValueSet() {

}

func (*TimeBase) ValueGet() {

}

func (*TimeBase) SearchSQLGet() {

}

func (*TimeBase) EditFieldRender() {

}

func (*TimeBase) EditFieldValueGet() {

}

func (*TimeBase) SearchFieldRender() {

}

func (*TimeBase) StatsFieldParameterBuild() {

}
