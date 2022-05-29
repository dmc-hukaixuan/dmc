package template

import (
	model "dmc/kernel/model/ticket"
)

type Dropdown struct{}

func (*Dropdown) TemplateEditRender() model.FeildData {
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

func (*Dropdown) ValueSet() {

}

func (*Dropdown) ValueGet() {

}

func (*Dropdown) SearchSQLGet() {

}

func (*Dropdown) EditFieldRender() {

}

func (*Dropdown) EditFieldValueGet() {

}

func (*Dropdown) SearchFieldRender() {

}

func (*Dropdown) StatsFieldParameterBuild() {

}
