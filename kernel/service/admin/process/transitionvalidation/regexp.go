package transitionvalidation

type Regexp struct{}

func (*Regexp) Validate(fieldType string, fieldValue interface{}, ticketFieldValue interface{}) bool {
	match := false

	return match
}
