package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"ruoyi-go/config"
	"sync"
	"time"
)

var once sync.Once

// 定义数据库连接结构体
type connect struct {
	primaryDB *gorm.DB
	currentDB *gorm.DB
}

// 设置一个全局变量保存数据库连接
var _connect *connect

func SqliteDb(param ...bool) *gorm.DB {
	if _connect == nil {
		once.Do(func() {
			connectSqlite()
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

func connectSqlite() {
	// 启用打印日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志等级: Silent、Error、Warn、Info
			Colorful:      false,       // 禁用彩色打印
		},
	)

	// sqlite数据库地址
	primary := "../sqlite-ry-go.db"

	// 打开主库连接
	primaryDB, err := gorm.Open(sqlite.Open(primary), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Printf("主数据库连接失败:", err)
		return
	}

	sqlDB, err := primaryDB.DB()
	if err != nil {
		log.Printf("获取从数据库SQL DB失败: %v", err)
		return
	}
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 初始化数据库连接结构体，默认使用主数据库
	_connect = &connect{
		primaryDB: primaryDB, // 默认连接到主数据库
	}
}
