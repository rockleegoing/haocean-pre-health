package db

import (
	"gorm.io/gorm"
	"ruoyi-go/config"
	"ruoyi-go/pkg/mysql"
	"ruoyi-go/pkg/postgres"
	"ruoyi-go/pkg/sqlite"
)

func Exe(param ...bool) *gorm.DB {
	dbtype := config.Database.Primary.Type
	switch dbtype {
	case "mysql":
		return mysql.MysqlDb(param...)
	case "sqlite":
		return sqlite.SqliteDb(param...)
	case "postgres":
		return postgres.PostgresDb(param...)
	}
	return mysql.MysqlDb(param...)
}
