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
	return &ServiceContext{
		Config:     c,
		UserModel2: model.NewUserModel2(db),
	}
}
