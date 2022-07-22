package transitionvalidation

type Equal struct{}

func (*Equal) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
