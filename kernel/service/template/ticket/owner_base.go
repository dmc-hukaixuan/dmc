package ticket

import (
	model "dmc/kernel/model/ticket"
)

type OwnerBase struct{}

func (*OwnerBase) TemplateEditRender(field string, ) model.FeildData {
	var fieldData model.FeildData
	fieldData = model.FeildData{
		Name:                 "Title",
		Default:              "",
		FieldType:            "text",
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

func (*OwnerBase) ValueSet() {

}

func (*OwnerBase) ValueGet() {

}

func (*OwnerBase) SearchSQLGet() {

}

func (*OwnerBase) EditFieldRender() {

}

func (*OwnerBase) EditFieldValueGet() {

}

func (*OwnerBase) SearchFieldRender() {

}

func (*OwnerBase) StatsFieldParameterBuild() {

}
