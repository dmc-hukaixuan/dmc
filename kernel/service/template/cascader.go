package template

import (
	model "dmc/kernel/model/ticket"
)

type Cascader struct{}

func (*Cascader) TemplateEditRender() model.FeildData {
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

func (*Cascader) ValueSet() {

}

func (*Cascader) ValueGet() {

}

func (*Cascader) SearchSQLGet() {

}

func (*Cascader) EditFieldRender() {

}

func (*Cascader) EditFieldValueGet() {

}

func (*Cascader) SearchFieldRender() {

}

func (*Cascader) StatsFieldParameterBuild() {

}
