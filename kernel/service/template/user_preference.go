package template

import (
	model "dmc/kernel/model/ticket"
)

type UserPreference struct{}

func (*UserPreference) TemplateEditRender() model.FeildData {
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

func (*UserPreference) ValueSet() {

}

func (*UserPreference) ValueGet() {

}

func (*UserPreference) SearchSQLGet() {

}

func (*UserPreference) EditFieldRender() {

}

func (*UserPreference) EditFieldValueGet() {

}

func (*UserPreference) SearchFieldRender() {

}

func (*UserPreference) StatsFieldParameterBuild() {

}
