package transitionvalidation

type IsOneOf struct{}

func (*IsOneOf) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
