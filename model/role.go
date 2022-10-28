package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
)

type Role struct {
	common.DBBase

	form.Role
}

func (Role) TableName() string {
	return "roles"
}

func (r Role) List(limit, offset int) ([]Role, error) {
	var roles []Role
	err := db.MysqlClient.Table(r.TableName()).Offset(offset).Limit(limit).Find(&roles).Error
	return roles, err
}

func (r *Role) GetByID(id string) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Find(&r).Error
}

func (r *Role) GetByName(name string) error {
	return db.MysqlClient.Table(r.TableName()).Where("role_name = ?", name).Find(&r).Error
}

func (r *Role) Update() error {
	return db.MysqlClient.Table(r.TableName()).Save(&r).Error
}

func (r *Role) Create() error {
	return db.MysqlClient.Table(r.TableName()).Create(&r).Error
}

func (r *Role) Del(id string) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Delete(&r).Error
}
