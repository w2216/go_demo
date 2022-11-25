package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// 用户表
type Admin struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        string
	Password     string
	Age          uint8
	Birthday     string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 登录记录表
type LoginLog struct {
	ID        uint `gorm:"primaryKey"`
	Aid       uint
	CreatedAt int
}

func init() {
	db := Db
	_ = db.AutoMigrate(&Admin{})
	_ = db.AutoMigrate(&LoginLog{})
}
