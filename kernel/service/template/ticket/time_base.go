package ticket

import (
	model "dmc/kernel/model/ticket"
)

type TimeBase struct{}

func (*TimeBase) TemplateEditRender() model.FeildData {
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
