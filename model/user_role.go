package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
)

type UserRole struct {
	common.DBBase

	form.RelationReq
}

func (UserRole) TableName() string {
	return "user_roles"
}

func (u *UserRole) Update() error {
	return db.MysqlClient.Table(u.TableName()).Save(u).Error
}

func (u *UserRole) Create() error {
	return db.MysqlClient.Table(u.TableName()).Create(u).Error
}

func (u *UserRole) GetByUserID(uid uint) error {
	return db.MysqlClient.Table(u.TableName()).Where("user_id = ?", uid).Find(u).Error
}

func (u UserRole) List(limit, offset int) ([]UserRole, error) {
	var list []UserRole
	err := db.MysqlClient.Table(u.TableName()).Offset(offset).Limit(limit).Find(&list).Error
	return list, err
}
