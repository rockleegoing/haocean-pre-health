package db

import (
	"gorm.io/gorm"
	"haocean/health-enforcement/config"
	"haocean/health-enforcement/pkg/mysql"
	"haocean/health-enforcement/pkg/postgres"
	"haocean/health-enforcement/pkg/sqlite"
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
