package ticket

import (
    model "dmc/kernel/model/admin"
    "encoding/json"

    //model "dmc/kernel/model/ticket"
    "dmc/kernel/service/admin/ticket"
    //"encoding/json"
)

type Priority struct{}

func (*Priority) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
    var fieldData model.FieldData
    priorityList := ticket.PriorityList(1)
    // get template or
    if fieldObject.FieldKey != "" {
        //json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
        fieldData.Options = priorityList
        return fieldData
    } else {
        fieldData = model.FieldData{
            Name:          "priority",
            Default:       "",
            FieldType:     "dropdown",
            Label:         "Priority",
            Placeholder:   "ticket service",
            Display:       2,
            Options:       priorityList,
            PromptCode:    2,
            PromptMessage: "",
            HintMessage:   "",
            Width:         4,
        }
    }

    return fieldData
}

func (*Priority) ValueSet() {

}

func (*Priority) ValueGet() {

}

func (*Priority) SearchSQLGet() {

}

func (*Priority) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {

    var perference_data model.FieldData
    priorityList := ticket.PriorityList(1)
    json.Unmarshal([]byte(fieldObject.FieldPreference), &perference_data)
    perference_data.Options = priorityList

    // 排除或者可选某个选项
    return perference_data
}

func (*Priority) EditFieldValueGet() {

}

func (*Priority) SearchFieldRender() model.FieldData {
    priorityList := ticket.PriorityList(1)
    fieldData := model.FieldData{
        Name:        "priority",
        Default:     "",
        FieldType:   "dropdown",
        Label:       "Priority",
        Placeholder: "",
        Options:     priorityList,
        Display:     1,
    }
    return fieldData
}

func (*Priority) StatsFieldParameterBuild() {

}
