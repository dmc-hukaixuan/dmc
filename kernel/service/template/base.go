package template

type DynamicFieldObject interface {
	TicketNumberBuild() string
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
		return &BaseTime{}
	case "date":
		return &BaseTime{}
	case "dropdown":
		return &Dropdown{}
	case "radio":
		return &Checkbox{}
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
		return &Date{}
	}
}
