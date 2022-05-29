package template

import (
	model "dmc/kernel/model/ticket"
)

type Tree struct{}

func (*Tree) TemplateEditRender() model.FeildData {
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

func (*Tree) ValueSet() {

}

func (*Tree) ValueGet() {

}

func (*Tree) SearchSQLGet() {

}

func (*Tree) EditFieldRender() {

}

func (*Tree) EditFieldValueGet() {

}

func (*Tree) SearchFieldRender() {

}

func (*Tree) StatsFieldParameterBuild() {

}
