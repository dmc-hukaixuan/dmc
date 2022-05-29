package template

import (
	model "dmc/kernel/model/ticket"
)

type Attachment struct{}

func (*Attachment) TemplateEditRender() model.FeildData {
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

func (*Attachment) ValueSet() {

}

func (*Attachment) ValueGet() {

}

func (*Attachment) SearchSQLGet() {

}

func (*Attachment) EditFieldRender() {

}

func (*Attachment) EditFieldValueGet() {

}

func (*Attachment) SearchFieldRender() {

}

func (*Attachment) StatsFieldParameterBuild() {

}
