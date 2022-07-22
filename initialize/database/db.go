package database

import (
	"dmc/config/database"
	"fmt"

	"gorm.io/gorm"
)

func Gorm(dbtype string) *gorm.DB {
	fmt.Println("dbtype :", database.GetConfig().DbType)
	switch dbtype {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "mysqlReport":
		return GormMysqlReport()
	default:
		return GormMysql()
	}
}
