package ticket

import (
	model "dmc/kernel/model/ticket"
)

type Priority struct{}

func (*Priority) TemplateEditRender() model.FeildData {
	var fieldData model.FeildData
	fieldData = model.FeildData{
		Name:                 "",
		Default:              "",
		FieldType:            "",
		Label:                "",
		Placeholder:          "",
		Display:              3,
		Impacts:              []string{},
		DependsOn:            []string{},
		PromptCode:           2,
		PromptMessage:        "",
		AutoComplete:         true,
		Options:              map[string]string{},
		OptionsType:          "",
		OptionsValueComments: map[string]string{},
		HintMessage:          "",
		HintType:             2,
		Width:                4,
		RegexError:           "",
		Regex:                "",
	}

	return fieldData
}

func (*Priority) ValueSet() {

}

func (*Priority) ValueGet() {

}

func (*Priority) SearchSQLGet() {

}

func (*Priority) EditFieldRender() {

}

func (*Priority) EditFieldValueGet() {

}

func (*Priority) SearchFieldRender() {

}

func (*Priority) StatsFieldParameterBuild() {

}
