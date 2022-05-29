package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// GormPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormPgSql() *gorm.DB {
	return db
}
