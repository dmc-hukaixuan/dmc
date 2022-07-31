package ticket

import (
	model "dmc/kernel/model/admin"
	//model "dmc/kernel/model/ticket"
	"dmc/kernel/service/admin/ticket"
	//"encoding/json"
)

type SLA struct{}

func (*SLA) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	slaList := ticket.SLAList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = slaList
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "source",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "source",
			Placeholder:   "ticket source",
			Display:       1,
			Options:       slaList,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}

	return fieldData
}

func (*SLA) ValueSet() {

}

func (*SLA) ValueGet() {

}

func (*SLA) SearchSQLGet() {

}

func (*SLA) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
	var fieldData model.FieldData

	fieldData = model.FieldData{
		Name:          "title",
		Default:       "",
		FieldType:     "text",
		Label:         fieldObject.FieldKey,
		Placeholder:   "",
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

func (*SLA) EditFieldValueGet() {

}

func (*SLA) SearchFieldRender() model.FieldData {
	slaList := ticket.SLAList(0)
	fieldData := model.FieldData{
		Name:        "service",
		Default:     "",
		FieldType:   "Tree",
		Label:       "Service",
		Placeholder: "",
		Options:     slaList,
		Display:     1,
	}
	return fieldData
}

func (*SLA) StatsFieldParameterBuild() {

}
