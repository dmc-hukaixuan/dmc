package ticket

import (
    model "dmc/kernel/model/admin"
    "encoding/json"

    //model "dmc/kernel/model/ticket"
    "dmc/kernel/service/admin/ticket"
    //"encoding/json"
)

type TicketSource struct{}

func (*TicketSource) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
    var fieldData model.FieldData
    // get ticket source list
    // ticket create be which way
    sourcelist := ticket.SourceList(1)
    // get template or
    if fieldObject.FieldKey != "" {
        //json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
        fieldData.Options = sourcelist
        return fieldData
    } else {
        fieldData = model.FieldData{
            Name:          "source",
            Default:       "",
            FieldType:     "dropdown",
            Label:         "source",
            Placeholder:   "ticket source",
            Display:       2,
            Options:       sourcelist,
            PromptCode:    2,
            PromptMessage: "",
            HintMessage:   "",
            Width:         4,
        }
    }

    return fieldData
}

func (*TicketSource) ValueSet() {

}

func (*TicketSource) ValueGet() {

}

func (*TicketSource) SearchSQLGet() {

}

func (*TicketSource) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
    var perference_data model.FieldData
    sourcelist := ticket.SourceList(1)
    json.Unmarshal([]byte(fieldObject.FieldPreference), &perference_data)
    perference_data.Options = sourcelist

    // 排除或者可选某个选项
    return perference_data
}

func (*TicketSource) EditFieldValueGet() {

}

func (*TicketSource) SearchFieldRender() model.FieldData {
    sourceList := ticket.SourceList(1)
    fieldData := model.FieldData{
        Name:        "source_id",
        Default:     "",
        FieldType:   "dropdown",
        Label:       "Source",
        Placeholder: "",
        Options:     sourceList,
        Display:     1,
    }
    return fieldData
}

func (*TicketSource) StatsFieldParameterBuild() {

}
