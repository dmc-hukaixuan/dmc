package router

import (
	"dmc/kernel/router/system"
)

type RouterGroup struct {
	SystemApiGroup system.ApiGroup
}

var RouterGroupApp = new(RouterGroup)
