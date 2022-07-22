package transitionaction

type LinkAdd struct{}

func (*LinkAdd) Run(ticketID int64, config map[string]interface{}) bool {
	success := true
	return success
}
