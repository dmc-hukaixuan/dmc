package transitionvalidation

type IsNotOneOf struct{}

func (*IsNotOneOf) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
