package ticket

import (
	model "dmc/kernel/model/ticket"
)

type TicketSource struct{}

func (*TicketSource) TemplateEditRender() model.FeildData {
	var fieldData model.FeildData
	fieldData = model.FeildData{
		Name:                 "",
		Default:              "",
		FieldType:            "",
		Label:                "",
		Placeholder:          "",
		Display:              3,
		Impacts:              []string{},
		DependsOn:            []string{},
		PromptCode:           2,
		PromptMessage:        "",
		AutoComplete:         true,
		Options:              map[string]string{},
		OptionsType:          "",
		OptionsValueComments: map[string]string{},
		HintMessage:          "",
		HintType:             2,
		Width:                4,
		RegexError:           "",
		Regex:                "",
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
