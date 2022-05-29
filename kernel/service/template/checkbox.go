package template

import (
	model "dmc/kernel/model/ticket"
)

type Checkbox struct{}

func (*Checkbox) TemplateEditRender() model.FeildData {
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

func (*Checkbox) ValueSet() {

}

func (*Checkbox) ValueGet() {

}

func (*Checkbox) SearchSQLGet() {

}

func (*Checkbox) EditFieldRender() {

}

func (*Checkbox) EditFieldValueGet() {

}

func (*Checkbox) SearchFieldRender() {

}

func (*Checkbox) StatsFieldParameterBuild() {

}
