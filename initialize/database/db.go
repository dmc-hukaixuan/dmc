package database

import (
	"dmc/config/database"
	"fmt"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	fmt.Println("dbtype :", database.GetConfig().DbType)
	switch database.GetConfig().DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}
