package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	"dmc/kernel/service/admin/ticket"
	//"encoding/json"
)

type TicketSource struct{}

func (*TicketSource) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get ticket source list
	// ticket create be which way
	sourcelist := ticket.SourceList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = sourcelist
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "source",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "source",
			Placeholder:   "ticket source",
			Display:       2,
			Options:       sourcelist,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}

	return fieldData
}

func (*TicketSource) ValueSet() {

}

func (*TicketSource) ValueGet() {

}

func (*TicketSource) SearchSQLGet() {

}

func (*TicketSource) EditFieldRender() {

}

func (*TicketSource) EditFieldValueGet() {

}

func (*TicketSource) SearchFieldRender() {

}

func (*TicketSource) StatsFieldParameterBuild() {

}
