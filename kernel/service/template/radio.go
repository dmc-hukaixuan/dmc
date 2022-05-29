package template

import (
	model "dmc/kernel/model/ticket"
)

type Radio struct{}

func (*Radio) TemplateEditRender() model.FeildData {
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

func (*Radio) ValueSet() {

}

func (*Radio) ValueGet() {

}

func (*Radio) SearchSQLGet() {

}

func (*Radio) EditFieldRender() {

}

func (*Radio) EditFieldValueGet() {

}

func (*Radio) SearchFieldRender() {

}

func (*Radio) StatsFieldParameterBuild() {

}
