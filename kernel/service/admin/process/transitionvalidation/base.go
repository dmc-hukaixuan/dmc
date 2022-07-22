package transitionvalidation

type TransitionValidation interface {
	Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool
}

func Validation(compareType string) TransitionValidation {
	switch compareType {
	case "eq":
		return &Equal{}
	case "neq":
		return &NotEqual{}
	case "ioo":
		return &IsOneOf{}
	case "inoo":
		return &IsNotOneOf{}
	case "regex":
		return &Regexp{}
	// case "contains":
	// 	return &Contains{}
	case "lessthan":
		return &LessThan{}
	default:
		return &Equal{}
	}
}
