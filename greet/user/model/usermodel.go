package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)
var _ UserModel2 = (*customUserModel2)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		// FindOneByName(ctx context.Context, name string) (*User, error)
	}

	UserModel2 interface {
		Add(ctx context.Context, name, phone, password string) (*User, error)
		Edit(ctx context.Context, id int64, name, phone, password string) (*User, error)
		Del(ctx context.Context, id int64) (*User, error)
		List(ctx context.Context, id int64, name, phone, password string) (*[]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserModel2 struct {
		*defaultUserModel2
	}

	defaultUserModel2 struct {
		conn  *gorm.DB
		table string
	}
)

func newUserModel2(conn *gorm.DB) *defaultUserModel2 {
	return &defaultUserModel2{
		conn:  conn,
		table: "`user`",
	}
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func NewUserModel2(conn *gorm.DB) UserModel2 {
	return &customUserModel2{
		defaultUserModel2: newUserModel2(conn),
	}
}

func (m *defaultUserModel2) Add(ctx context.Context, name, phone, password string) (*User, error) {
	user := User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	// 新增
	res := m.conn.Table("user").Create(&user)

	// 指定字段新增
	// res := m.conn.Table("user").Select("Name", "Phone").Create(&user)

	// 批量新增
	//users := []User{
	//	{
	//		Name:     name,
	//		Phone:    phone,
	//		Password: password,
	//	},
	//	{
	//		Name:     name,
	//		Phone:    phone,
	//		Password: password,
	//	},
	//	{
	//		Name:     name,
	//		Phone:    phone,
	//		Password: password,
	//	},
	//}
	//res := m.conn.Debug().Table("user").CreateInBatches(users, 1)

	logx.Info(res.Error)
	logx.Info(res.RowsAffected)

	return &user, nil
}

func (m *defaultUserModel2) Edit(ctx context.Context, id int64, name, phone, password string) (*User, error) {
	user := User{
		Id:       id,
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	// 更新 或 新增
	res := m.conn.Debug().Table("user").Save(&user)
	// 更新指定字段
	//res := m.conn.Debug().Table("user").Where("id = ?", id).Update("phone", phone)
	// 更新多列
	//res := m.conn.Debug().Table("user").Where("id = ?", id).Updates(user)

	if res.Error != nil {
		logx.Info(res.Error)
		logx.Info(res.RowsAffected)
		return nil, res.Error
	}

	return &user, nil
}

func (m *defaultUserModel2) Del(ctx context.Context, id int64) (*User, error) {
	user := User{
		Id: id,
	}
	// 删除
	res := m.conn.Debug().Table("user").Delete(&user)
	// 主键 删除
	//res := m.conn.Debug().Table("user").Delete(&User{}, 10)
	// 根据条件 删除
	//res := m.conn.Debug().Table("user").Where("phone = ?", "13822222222").Delete(&User{})

	if res.Error != nil {
		logx.Info(res.Error)
		logx.Info(res.RowsAffected)
		return nil, res.Error
	}

	return &user, nil
}

func (m *defaultUserModel2) List(ctx context.Context, id int64, name, phone, password string) (*[]User, error) {
	var users []User

	// 查询
	//res := m.conn.Debug().Table("user").Find(&users)
	res := m.conn.Debug().Table("user").Limit(3).Find(&users)
	// 条件查询
	//res := m.conn.Debug().Table("user").Select("id", "name", "password").Order("id desc").Find(&users)
	//res := m.conn.Debug().Table("user").Where("phone = ? ","18326117467").Find(&users)
	//res := m.conn.Debug().Table("user").Where("phone in ?", []string{"18326117467"}).Find(&users)

	if res.Error != nil {
		logx.Info(res.Error)
		logx.Info(res.RowsAffected)
		return nil, res.Error
	}

	return &users, nil
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {

	if user.Id == 23 {
		user.Phone = "111111"
	}
	return
}
