package template

import (
	model "dmc/kernel/model/ticket"
)

type Textarea struct{}

func (*Textarea) TemplateEditRender() model.FeildData {
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

func (*Textarea) ValueSet() {

}

func (*Textarea) ValueGet() {

}

func (*Textarea) SearchSQLGet() {

}

func (*Textarea) EditFieldRender() {

}

func (*Textarea) EditFieldValueGet() {

}

func (*Textarea) SearchFieldRender() {

}

func (*Textarea) StatsFieldParameterBuild() {

}
