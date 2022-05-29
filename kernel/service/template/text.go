package template

import (
	model "dmc/kernel/model/ticket"
)

type Text struct{}

func (*Text) TemplateEditRender() model.FeildData {
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

func (*Text) ValueSet() {

}

func (*Text) ValueGet() {

}

func (*Text) SearchSQLGet() {

}

func (*Text) EditFieldRender() {

}

func (*Text) EditFieldValueGet() {

}

func (*Text) SearchFieldRender() {

}

func (*Text) StatsFieldParameterBuild() {

}
