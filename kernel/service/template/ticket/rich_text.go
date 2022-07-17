package ticket

import (
    model "dmc/kernel/model/admin"
    //model "dmc/kernel/model/ticket"
    //"encoding/json"
)

type Richtext struct{}

func (*Richtext) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
    var fieldData model.FieldData
    // get template or
    if fieldObject.FieldKey != "" {
        //json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
        return fieldData
    } else {
        fieldData = model.FieldData{
            Name:          "body",
            Default:       "",
            FieldType:     "richText",
            Label:         "Content",
            Placeholder:   "",
            Display:       2,
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

func (*Richtext) ValueSet() {

}

func (*Richtext) ValueGet() {

}

func (*Richtext) SearchSQLGet() {

}

func (*Richtext) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
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

func (*Richtext) EditFieldValueGet() {

}

func (*Richtext) SearchFieldRender() {

}

func (*Richtext) StatsFieldParameterBuild() {

}
