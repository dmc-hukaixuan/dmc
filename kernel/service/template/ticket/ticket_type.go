package ticket

import (
	model "dmc/kernel/model/ticket"
)

type TicketType struct{}

func (*TicketType) TemplateEditRender() model.FeildData {
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
