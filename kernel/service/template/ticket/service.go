package ticket

import model "dmc/kernel/model/ticket"

type Service struct{}

func (*Service) TemplateEditRender() model.FeildData {
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

func (*Service) ValueSet() {

}

func (*Service) ValueGet() {

}

func (*Service) SearchSQLGet() {

}

func (*Service) EditFieldRender() {

}

func (*Service) EditFieldValueGet() {

}

func (*Service) SearchFieldRender() {

}

func (*Service) StatsFieldParameterBuild() {

}
