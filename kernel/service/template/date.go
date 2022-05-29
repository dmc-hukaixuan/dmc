package template

import (
	model "dmc/kernel/model/ticket"
)

type Date struct{}

func (*Date) TemplateEditRender() model.FeildData {
	var fieldData model.FeildData
	fieldData = model.FeildData{
		Name:                 "",
		Default:              "",
		FieldType:            "",
		Label:                "",
		Placeholder:          "",
		Display:              "",
		Impacts:              "",
		DependsOn:            "",
		PromptCode:           "",
		PromptMessage:        "",
		AutoComplete:         "",
		Options:              "",
		OptionsType:          "",
		OptionsValueComments: "",
		HintMessage:          "",
		HintType:             "",
		Width:                "",
		RegexError:           "",
		Regex:                "",
	}

	return fieldData
}

func (*Date) ValueSet() {

}

func (*Date) ValueGet() {

}

func (*Date) SearchSQLGet() {

}

func (*Date) EditFieldRender() {

}

func (*Date) EditFieldValueGet() {

}

func (*Date) SearchFieldRender() {

}

func (*Date) StatsFieldParameterBuild() {

}
