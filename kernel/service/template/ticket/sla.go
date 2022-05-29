package ticket

import (
	model "dmc/kernel/model/ticket"
)

type SLA struct{}

func (*SLA) TemplateEditRender() model.FeildData {
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

func (*SLA) ValueSet() {

}

func (*SLA) ValueGet() {

}

func (*SLA) SearchSQLGet() {

}

func (*SLA) EditFieldRender() {

}

func (*SLA) EditFieldValueGet() {

}

func (*SLA) SearchFieldRender() {

}

func (*SLA) StatsFieldParameterBuild() {

}
