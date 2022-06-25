package ticket

import (
	model "dmc/kernel/model/admin"
	// model "dmc/kernel/model/ticket"
	// "encoding/json"
)

type Title struct{}

/**/
func (*Title) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	// get template or
	if fieldObject.FieldKey != "" {
		//json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "title",
			Default:       "",
			FieldType:     "text",
			Label:         fieldObject.FieldKey,
			Placeholder:   "",
			Display:       fieldObject.Display,
			Impacts:       []string{},
			DependsOn:     []string{},
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			HintType:      2,
			Width:         4,
			RegexError:    "",
			Regex:         "",
		}
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
