package transitionaction

type TransitionAtion interface {
	Run(ticketID int64, config map[string]interface{}) bool
}

func TransitionAction(module string) TransitionAtion {
	switch module {
	case "ticketTemplate":
		return &TicketTemplate{}
	case "linkAdd":
		return &LinkAdd{}
	case "itsm":
		return &ConfigItem{}
	default:
		return &TicketTemplate{}
	}
}
