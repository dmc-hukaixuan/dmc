package template

import (
	model "dmc/kernel/model/ticket"
)

type Richtext struct{}

func (*Richtext) TemplateEditRender() model.FeildData {
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

func (*Richtext) ValueSet() {

}

func (*Richtext) ValueGet() {

}

func (*Richtext) SearchSQLGet() {

}

func (*Richtext) EditFieldRender() {

}

func (*Richtext) EditFieldValueGet() {

}

func (*Richtext) SearchFieldRender() {

}

func (*Richtext) StatsFieldParameterBuild() {

}
