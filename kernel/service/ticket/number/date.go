package number

type Date struct{}

func (*Date) TicketNumberBuild() string {
	// get current ticket counter
	return "20999"
}
