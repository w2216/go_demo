package model

import (
	"database/sql"
	"fmt"
	"gin/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbErr error

var SqlDB *sql.DB
var SqlDBErr error

const (
	USER_NAME = "root"
	PASS_WORD = "a88888888"
	HOST      = "sddphp.cn"
	PORT      = "19906"
	DATABASE  = "gorm"
	CHARSET   = "utf8"
)

// 初始化链接
func init() {
	slowLogger := logger.New(
		//设置Logger
		NewMyWriter(),
		logger.Config{
			//慢SQL阈值
			// SlowThreshold: time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel: logger.Info,
		},
	)

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	// 打开连接失败
	Db, DbErr = gorm.Open(mysql.New(mysql.Config{
		DSN: dbDSN, // DSN data source name
		// DefaultStringSize:         256,   // string 类型字段的默认长度
		// DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		// DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		// DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		// SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: slowLogger,
	})
	if DbErr != nil {
		panic("数据库链接失败1: " + DbErr.Error())
	}

	SqlDB, SqlDBErr = Db.DB()

	if SqlDBErr != nil {
		panic("数据库链接失败2: " + SqlDBErr.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDB.SetConnMaxLifetime(time.Hour)

	// ping
	_ = SqlDB.Ping()

}

// 定义自己的Writer
type MyWriter struct {
	mlog *logrus.Logger
}

// 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	log := logrus.New()
	//配置logrus
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	src, _ := middleware.LogSrc()
	log.Out = src

	return &MyWriter{mlog: log}
}
