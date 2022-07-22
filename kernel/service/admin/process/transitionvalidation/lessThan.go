package transitionvalidation

type LessThan struct{}

func (*LessThan) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
