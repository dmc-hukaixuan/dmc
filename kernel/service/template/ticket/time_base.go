package ticket

import (
    model "dmc/kernel/model/admin"
    //model "dmc/kernel/model/ticket"
    //"encoding/json"
)

type TimeBase struct{}

func (*TimeBase) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
    var fieldData model.FieldData
    // get template or
    if fieldObject.FieldKey != "" {
        //json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
        return fieldData
    } else {
        // incident start time
        // incident end time
        fieldData = model.FieldData{
            Name:          "title",
            Default:       "",
            FieldType:     "text",
            Label:         fieldObject.FieldKey,
            Placeholder:   "",
            Display:       fieldObject.Display,
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

func (*TimeBase) ValueSet() {

}

func (*TimeBase) ValueGet() {

}

func (*TimeBase) SearchSQLGet() {

}

func (*TimeBase) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
    var fieldData model.FieldData

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
    return fieldData
}

func (*TimeBase) EditFieldValueGet() {

}

func (*TimeBase) SearchFieldRender() model.FieldData {
	fieldData := model.FieldData{
		Name:        "source_id",
		Default:     "",
		FieldType:   "dropdown",
		Label:       "Source",
		Placeholder: "",
		Display:     1,
	}
	return fieldData
}

func (*TimeBase) StatsFieldParameterBuild() {

}
