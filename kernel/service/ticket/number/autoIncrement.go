package number

type AutoIncrement struct {
}

func (*AutoIncrement) TicketNumberBuild() string {
	return "20999"
}
