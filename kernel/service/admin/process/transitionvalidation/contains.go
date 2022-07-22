package transitionvalidation

type Contains struct{}

func (*Contains) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
