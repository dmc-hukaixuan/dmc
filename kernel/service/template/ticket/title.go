package ticket

import (
	model "dmc/kernel/model/ticket"
)

type Title struct{}

/**/
func (*Title) TemplateEditRender() model.FeildData {
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

func (*Title) ValueSet() {

}

func (*Title) ValueGet() {

}

func (*Title) SearchSQLGet() {

}

func (*Title) EditFieldRender() {

}

func (*Title) EditFieldValueGet() {

}

func (*Title) SearchFieldRender() {

}

func (*Title) StatsFieldParameterBuild() {

}
