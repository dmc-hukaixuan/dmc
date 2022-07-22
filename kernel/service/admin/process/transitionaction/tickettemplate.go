package transitionaction

type TicketTemplate struct{}

func (*TicketTemplate) Run(ticketID int64, config map[string]interface{}) bool {
	success := true
	return success
}
