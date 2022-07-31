package ticket

import (
	model "dmc/kernel/model/admin"
	"encoding/json"

	//model "dmc/kernel/model/ticket"
	"dmc/kernel/service/admin/ticket"
	//"encoding/json"
)

type Service struct{}

func (*Service) TemplateEditRender(fieldType string, fieldObject model.TemplateField) model.FieldData {
	var fieldData model.FieldData
	serviceList := ticket.ServiceList(1)
	// get template or
	if fieldObject.FieldKey != "" {
		//	json.Unmarshal([]byte(fieldObject.FieldPreference), &fieldData)
		fieldData.Options = serviceList
		return fieldData
	} else {
		fieldData = model.FieldData{
			Name:          "service",
			Default:       "",
			FieldType:     "dropdown",
			Label:         "Service",
			Placeholder:   "ticket service",
			Display:       1,
			Options:       serviceList,
			PromptCode:    2,
			PromptMessage: "",
			HintMessage:   "",
			Width:         4,
		}
	}

	return fieldData
}

func (*Service) ValueSet() {

}

func (*Service) ValueGet() {

}

func (*Service) SearchSQLGet() {

}

func (*Service) EditFieldRender(fieldType string, fieldObject model.TemplateField, value interface{}) model.FieldData {
	var perference_data model.FieldData
	json.Unmarshal([]byte(fieldObject.FieldPreference), &perference_data)
	serviceList := ticket.ServiceList(1)
	perference_data.Options = serviceList
	// fieldData := model.FieldData{
	// 	Name:          "title",
	// 	Default:       "",
	// 	FieldType:     "text",
	// 	Label:         fieldObject.FieldKey,
	// 	Placeholder:   "",
	// 	Options:       serviceList,
	// 	Display:       fieldObject.Display,
	// 	Impacts:       []string{},
	// 	DependsOn:     []string{},
	// 	PromptCode:    2,
	// 	PromptMessage: "",
	// 	HintMessage:   "",
	// 	HintType:      2,
	// 	Width:         4,
	// 	RegexError:    "",
	// 	Regex:         "",
	// }
	return perference_data
}

func (*Service) EditFieldValueGet() {

}

func (*Service) SearchFieldRender() model.FieldData {
	serviceList := ticket.ServiceList(1)
    fieldData := model.FieldData{
        Name:        "service",
        Default:     "",
        FieldType:   "Tree",
        Label:       "Service",
        Placeholder: "",
        Options:     serviceList,
        Display:     1,
    }
    return fieldData
}

func (*Service) StatsFieldParameterBuild() {

}
