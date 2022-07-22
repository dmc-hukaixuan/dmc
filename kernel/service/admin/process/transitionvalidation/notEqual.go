package transitionvalidation

type NotEqual struct{}

func (*NotEqual) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
