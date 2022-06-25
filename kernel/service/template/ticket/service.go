package ticket

import (
	model "dmc/kernel/model/admin"
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

func (*Service) EditFieldRender() {

}

func (*Service) EditFieldValueGet() {

}

func (*Service) SearchFieldRender() {

}

func (*Service) StatsFieldParameterBuild() {

}
