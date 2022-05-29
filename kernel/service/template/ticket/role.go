package ticket

import model "dmc/kernel/model/ticket"

type Role struct{}

func (*Role) TemplateEditRender() model.FeildData {
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

func (*Role) ValueSet() {

}

func (*Role) ValueGet() {

}

func (*Role) SearchSQLGet() {

}

func (*Role) EditFieldRender() {

}

func (*Role) EditFieldValueGet() {

}

func (*Role) SearchFieldRender() {

}

func (*Role) StatsFieldParameterBuild() {

}
