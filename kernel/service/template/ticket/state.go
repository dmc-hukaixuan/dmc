package ticket

import model "dmc/kernel/model/ticket"

type TicketState struct{}

func (*TicketState) TemplateEditRender() model.FeildData {
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

func (*TicketState) ValueSet() {

}

func (*TicketState) ValueGet() {

}

func (*TicketState) SearchSQLGet() {

}

func (*TicketState) EditFieldRender() {

}

func (*TicketState) EditFieldValueGet() {

}

func (*TicketState) SearchFieldRender() {

}

func (*TicketState) StatsFieldParameterBuild() {

}
