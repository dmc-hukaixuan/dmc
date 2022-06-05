package v1

import (
	"dmc/api/v1/admin"
	"dmc/api/v1/system"
)

type APIGroup struct {
	Auth  system.APIGroup
	Admin admin.Admin
}

var APIGroupApp = new(APIGroup)
