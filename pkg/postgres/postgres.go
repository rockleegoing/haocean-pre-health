package postgres

import (
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

// 定义数据库连接结构体
type connect struct {
	primaryDB *gorm.DB
	currentDB *gorm.DB
}

// 设置一个全局变量保存数据库连接
var _connect *connect

func PostgresDb(param ...bool) *gorm.DB {
	if _connect == nil {
		once.Do(func() {
			connectPostgres()
		})
	}
	if len(param) == 0 {
		return _connect.primaryDB
	}
	if len(param) > 0 {
		return _connect.currentDB
	}
	return _connect.primaryDB
}

func connectPostgres() {

}
