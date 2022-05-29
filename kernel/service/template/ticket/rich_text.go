package ticket

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
