package transitionaction

type ConfigItem struct{}

func (*ConfigItem) Run(ticketID int64, config map[string]interface{}) bool {
	success := true
	return success
}
