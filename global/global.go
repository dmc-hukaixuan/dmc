package global

import (
	"dmc/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB
	GVA_DB_REPORT *gorm.DB
	GVA_LOG       *zap.Logger
	CONFIG        config.Server
)
