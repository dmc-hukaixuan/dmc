package template

import model "dmc/kernel/model/admin"

//model "dmc/kernel/model/ticket"

type DynamicFieldObject interface {
    TemplateEditRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField) model.FieldData
    EditFieldRender(fieldLabel string, fieldName string, DynamicFieldConfig *model.DynamicFieldConfig, FieldObject model.TemplateField, value interface{}) model.FieldData

    ValueSet(fieldID int, object string, objectID int64, value interface{})
    ValueGet(fieldID int, object string, objectID int64) interface{}
}

func DynamicField(field_type string) DynamicFieldObject {
    // 这里应该获取系统配置中的配置，实现工单单号的生成
    switch field_type {
    case "attachment":
        return &Attachment{}
    case "cascader":
        return &Cascader{}
    case "chackbose":
        return &Checkbox{}
    case "datetime":
        return &BaseDate{}
    case "date":
        return &BaseDate{}
    case "dropdown":
        return &Dropdown{}
    case "Dropdown":
        return &Dropdown{}
    case "radio":
        return &Checkbox{}
    case "Radio":
        return &Radio{}
    case "text":
        return &Text{}
    case "tree":
        return &Tree{}
    case "textarea":
        return &Textarea{}
    case "userpreference":
        return &UserPreference{}
    // case "aliyun-oss":
    // 	return &AliyunOSS{}
    default:
        return &Text{}
    }
}
