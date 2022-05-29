package template

import (
	model "dmc/kernel/model/ticket"
)

type RolePreference struct{}

func (*RolePreference) TemplateEditRender() model.FeildData {
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

func (*RolePreference) ValueSet() {

}

func (*RolePreference) ValueGet() {

}

func (*RolePreference) SearchSQLGet() {

}

func (*RolePreference) EditFieldRender() {

}

func (*RolePreference) EditFieldValueGet() {

}

func (*RolePreference) SearchFieldRender() {

}

func (*RolePreference) StatsFieldParameterBuild() {

}
