package dynmaicfield

//
type DynamicFieldObject interface {
	ValueSet(key string) error
	ValueGet(key string) error
	EditFieldRender() error
	TemplateRender() error
	SearchFieldRender() error
	DisplayValueRender() error
}

// interface dynamci field method, retrun this object
func DynamicField(fieldType string) DynamicFieldObject {
	switch fieldType {
	case "radio":
		return &Radio{}
	case "date":
		return &Date{}
	case "date_time":
		return &DateTime{}
	case "cascader":
		return &Cascader{}
	case "checkbox":
		return &Checkbox{}
	case "dropdown":
		return &Dropdown{}
	case "text":
		return &Text{}
	case "textArea":
		return &Textarea{}
	case "richText":
		return &Richtext{}
	case "number":
		return &Number{}
	default:
		return &Text{}
	}
}
