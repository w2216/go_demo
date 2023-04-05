package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"user/internal/config"
	"user/model"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	UserModel2 model.UserModel2
}

func NewServiceContext(c config.Config) *ServiceContext {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	dsn := c.Mysql.DataSource
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// 连接池配置
	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// 设置空闲连接池的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// 设置到数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// 设置连接可以被重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return &ServiceContext{
		Config:     c,
		UserModel2: model.NewUserModel2(db),
	}
}
