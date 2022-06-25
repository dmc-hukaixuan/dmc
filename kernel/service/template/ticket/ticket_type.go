package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	"dmc/kernel/service/admin/ticket"
	//"encoding/json"
)

type TicketType struct{}

func (*TicketType) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get ticket type list
	typelist := ticket.TypeList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = typelist
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "type",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "Type",
			Placeholder:   "",
			Display:       2,
			Options:       typelist,
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

func (*TicketType) ValueSet() {

}

func (*TicketType) ValueGet() {

}

func (*TicketType) SearchSQLGet() {

}

func (*TicketType) EditFieldRender() {

}

func (*TicketType) EditFieldValueGet() {

}

func (*TicketType) SearchFieldRender() {

}

func (*TicketType) StatsFieldParameterBuild() {

}
