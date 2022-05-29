package template

import (
	model "dmc/kernel/model/ticket"
)

type DataTime struct{}

func (*DataTime) TemplateEditRender() model.FeildData {
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

func (*DataTime) ValueSet() {

}

func (*DataTime) ValueGet() {

}

func (*DataTime) SearchSQLGet() {

}

func (*DataTime) EditFieldRender() {

}

func (*DataTime) EditFieldValueGet() {

}

func (*DataTime) SearchFieldRender() {

}

func (*DataTime) StatsFieldParameterBuild() {

}
