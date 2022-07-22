package transitionvalidation

type Input struct{}

func (*Input) Validate(FieldName string, Condition string) bool {
	match := false
	return match
}
