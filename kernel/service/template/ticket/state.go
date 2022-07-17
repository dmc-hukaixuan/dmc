package ticket

import (
	model "dmc/kernel/model/admin"
	"encoding/json"
	"fmt"

	//model "dmc/kernel/model/ticket"
	"dmc/kernel/service/admin/ticket"
	//"encoding/json"
)

type TicketState struct{}

func (*TicketState) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	statelist := ticket.StateList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = statelist
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "state",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "State",
			Placeholder:   "",
			Display:       2,
			Options:       statelist,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}

	return fieldData
}

func (*TicketState) ValueSet() {

}

func (*TicketState) ValueGet() {

}

func (*TicketState) SearchSQLGet() {

}

func (*TicketState) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
	var fieldData1 model.FieldData
	statelist := ticket.StateList(1)
	json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData1)
	fmt.Println("FieldPreference ------------------ ", fieldData1, fieldObject.FieldPreference)
	fieldData := model.FieldData{
		Name:          "title",
		Default:       "",
		FieldType:     "text",
		Label:         fieldObject.FieldKey,
		Placeholder:   "",
		Options:       statelist,
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
	return fieldData
}

func (*TicketState) EditFieldValueGet() {

}

func (*TicketState) SearchFieldRender() {

}

func (*TicketState) StatsFieldParameterBuild() {

}
