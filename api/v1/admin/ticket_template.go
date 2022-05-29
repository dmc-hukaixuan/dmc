package admin

import (
	"dmc/kernel/model/common/request"
	"dmc/kernel/model/common/response"
	model "dmc/kernel/model/ticket"
	system "dmc/kernel/service/admin/ticket"
	ticketField "dmc/kernel/service/template/ticket"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TicketTemplateApi struct {
	// BaseController
}

// ticket template base
func (p *TicketTemplateApi) Base(c *gin.Context) {

	var sd request.SubActionData
	_ = c.ShouldBindJSON(&sd)
	fmt.Println("sd ", sd.SubAction)
	if sd.SubAction == "edit" {
		templateID, _ := sd.Data["templateID"]
		fmt.Println(" templateID", templateID)
		TicketTemplateEdit("", c)
	}
}

type TemplateData struct {
	ID              int                     `json:"id"`
	Name            string                  `json:"name"`
	Web             string                  `json:"web"`
	Mobile          string                  `json:"mobile"`
	Valid_id        string                  `json:"valid_id"`
	Describe        string                  `json:"describe"`
	Icon            string                  `json:"icon"`
	Color           string                  `json:"color"`
	Roles           []int                   `json:"roles"`
	Type            string                  `json:"type"`
	FilterCondition *map[string]interface{} `json:"fieldOrder"`
	FieldOrder      []string                `json:"fieldOrder"`
	FieldData       *model.FeildData        `json:"fiedlData"`
	DisplayType     string                  `json:"display_type"`
}

// get ticket template data ,for add or edit a template
func TicketTemplateEdit(templateID string, c *gin.Context) {
	// get ticket template data

	// template base field

	// ticket base field
	baseFeild := map[string]*model.FeildData{
		"name": &model.FeildData{
			Name:      "name",
			Default:   "",
			FieldType: "text",
			Label:     "Process name",
			Display:   2,
		},
		"description": &model.FeildData{
			Name:      "description",
			Default:   "",
			FieldType: "textarea",
			Label:     "Description",
			Display:   1,
		},
		"valid": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label:   "Description",
			Display: 1,
		},
		"showLocation": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"mobile": "Show Mobile",
				"web":    "Show Web",
			},
			Label:   "Description",
			Display: 1,
		},
		"displayType": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "show form, edit field",
				"2": "not show form",
				"3": "open dailog",
			},
			Label:   "Description",
			Display: 1,
		},
		"templateType": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"create": "Create Normal Ticket",
				"const":  "Ticket Detail Show",
				"deal":   "Processing Ticket",
			},
			Label:   "Description",
			Display: 1,
		},
		"roles": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label:   "Description",
			Display: 1,
		},
		"color": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label:   "Description",
			Display: 1,
		},
		"icon": &model.FeildData{
			Name:      "valid",
			Default:   "",
			FieldType: "dropdown",
			Options: map[string]string{
				"1": "Valid",
				"2": "Invalid",
			},
			Label:   "Description",
			Display: 1,
		},
	}
	fmt.Println(" baseFeild", baseFeild)
	// dynamic field
	allticketField()

	response.SuccessWithDetailed(gin.H{
		"base": baseFeild,
		//	"filter":         allticketField(),
		"template":       "",
		"ticketRequired": "",
	}, "获取成功", c)

}

// get ticket template data ,for add or edit a template
func (p *TicketTemplateApi) TicketTemplateGet(c *gin.Context) {
	// get ticket template

	// base field

	// dynamic field
	// DynamicFieldList("Ticket")
	response.SuccessWithDetailed(gin.H{
		"baseField":           [...]string{"name", "description", "processState", "processType"},
		"filterField":         "",
		"templateField":       "",
		"ticketRequiredField": "",
		"":                    "",
	}, "获取成功", c)
}

// add copy etc...
func (p *TicketTemplateApi) TicketTemplateAdd(c *gin.Context) {
	// parse template data

}

// ticket template
func ticketTemplateUpdate() {

}

// remove ticket template
func ticketTemplateDelete() {

}

func ticketTemplateList() {

}

//
func allticketField() interface{} {
	pl := system.PriorityList(1)

	fmt.Println("pl ", pl)
	// ticketFieldObject := ticketField.TicketStartandField()
	templateField := []string{"title", "type"}
	for _, v := range templateField {
		fmt.Println("v ", v)
		fieldObject := ticketField.TicketStartandField(v).TemplateEditRender(v,)

		fmt.Println("fieldInfo ", fieldObject)
	}

	field := map[string]*model.FeildData{
		"title": &model.FeildData{
			Name:        "title",
			Default:     "",
			FieldType:   "text",
			Label:       "Title",
			Placeholder: "",
			Display:     1,
			Regex:       "",
			RegexError:  "",
			Width:       2,
		},
		"subject": &model.FeildData{
			Name:        "title",
			Default:     "",
			FieldType:   "text",
			Placeholder: "",
			Label:       "Title",
			Display:     1,
			Regex:       "",
			RegexError:  "",
			Width:       16,
		},
		"type": &model.FeildData{
			Name:        "title",
			Default:     "",
			FieldType:   "text",
			Label:       "Title",
			Display:     1,
			OptionsType: "",
			Options:     system.StateList(1),
			Regex:       "",
			RegexError:  "",
			Width:       16,
		},
		"state":           &model.FeildData{},
		"role":            &model.FeildData{},
		"source":          &model.FeildData{},
		"userid":          &model.FeildData{},
		"department":      &model.FeildData{},
		"lock":            &model.FeildData{},
		"priority":        &model.FeildData{},
		"service":         &model.FeildData{},
		"sla":             &model.FeildData{},
		"real_start_time": &model.FeildData{},
		"real_end_time":   &model.FeildData{},
		"customerid":      &model.FeildData{},
	}

	// get dynamic field list

	return field
}
