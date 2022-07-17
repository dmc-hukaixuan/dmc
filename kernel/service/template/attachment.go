package template

import (
    model "dmc/kernel/model/admin"
    model_dynamicField "dmc/kernel/model/dynamicField"
    service "dmc/kernel/service/ticket/dynamicField"
    // model "dmc/kernel/model/ticket"
)

type Attachment struct{}

func (*Attachment) TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, fieldObject model.TemplateField) model.FieldData {
    var fieldData model.FieldData

    // get template or
    if fieldObject.FieldKey != "" {

        return fieldData
    } else {
        fieldData = model.FieldData{
            Name:          fieldName,
            Default:       "",
            FieldType:     "attachment",
            Label:         fieldLabel,
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

func (*Attachment) ValueSet(fieldID int, object string, objectID int64, value interface{}) {
    values := []model_dynamicField.DynamicFieldValue{}
    values = append(values, model_dynamicField.DynamicFieldValue{
        FieldID:   fieldID,
        ObjectID:  objectID,
        ValueText: value,
    })
    service.ValueSet(fieldID, objectID, values)
}

func (*Attachment) ValueGet(fieldID int, object string, objectID int64) interface{} {
    values := service.ValueGet(fieldID, objectID)
    return values[0].ValueText
}

func (*Attachment) SearchSQLGet() {

}

func (*Attachment) EditFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField, value interface{}) model.FieldData {
    var fieldData model.FieldData

    fieldData = model.FieldData{
        Name:          "title",
        Default:       "",
        FieldType:     "cascader",
        Label:         FieldObject.FieldKey,
        Placeholder:   "",
        Display:       FieldObject.Display,
        Impacts:       []string{},
        DependsOn:     []string{},
        PromptCode:    2,
        PromptMessage: "",
        HintMessage:   "",
        HintType:      2,
        Width:         24,
        RegexError:    "",
        Regex:         "",
    }
    return fieldData
}

func (*Attachment) EditFieldValueGet() {

}

func (*Attachment) SearchFieldRender() {

}

func (*Attachment) StatsFieldParameterBuild() {

}
