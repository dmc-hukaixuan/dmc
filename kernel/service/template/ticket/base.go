package ticket

import (
    model "dmc/kernel/model/admin"
    //model "dmc/kernel/model/ticket"
)

type TF interface {
    TemplateEditRender(fieldtype string, FieldObject model.TemplateField) model.FieldData
    SearchSQLGet()
    EditFieldRender(fieldtype string, FieldObject model.TemplateField, value interface{}) model.FieldData
    SearchFieldRender() model.FieldData
}

/*
    interface
    @params : fieldType
    @return : each field type

use method:

*/
func TicketStartandField(fieldtype string) TF {
    switch fieldtype {
    case "user":
        return &OwnerBase{}
    case "owner":
        return &OwnerBase{}
    case "priority":
        return &Priority{}
    case "richtext":
        return &Richtext{}
    case "service":
        return &Service{}
    case "sla":
        return &SLA{}
    case "state":
        return &TicketState{}
    case "source":
        return &TicketSource{}
    case "type":
        return &TicketType{}
    case "time":
        return &TimeBase{}
    case "title":
        return &Title{}
    case "role":
        return &Role{}

    // case "aliyun-oss":
    // 	return &AliyunOSS{}
    default:
        return &Title{}
    }
}
