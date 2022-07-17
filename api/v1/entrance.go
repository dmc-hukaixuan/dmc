package v1

import (
    "dmc/api/v1/admin"
    "dmc/api/v1/system"
    "dmc/api/v1/ticket"
)

type APIGroup struct {
    Auth   system.APIGroup
    Admin  admin.Admin
    Ticket ticket.TicketAPI
}

var APIGroupApp = new(APIGroup)
