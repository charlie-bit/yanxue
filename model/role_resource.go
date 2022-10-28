package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
)

type RoleResourceColumns struct {
	common.DBBase

	RoleId     uint // 角色id
	ResourceId uint // 资源id
}

func (RoleResourceColumns) TableName() string {
	return "role_resource_columns"
}

func (r *RoleResourceColumns) Create() error {
	return db.MysqlClient.Table(r.TableName()).Create(&r).Error
}

func (r *RoleResourceColumns) Del(id uint) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Delete(&r).Error
}

func (r *RoleResourceColumns) GetRRID(rid, rsid uint) error {
	return db.MysqlClient.Table(r.TableName()).Where("role_id = ? and resource_id = ?", rid, rsid).Find(&r).Error
}

func (r *RoleResourceColumns) GetByID(id string) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Find(&r).Error
}

func (r *RoleResourceColumns) GetByRID(id string) ([]RoleResourceColumns, error) {
	var list []RoleResourceColumns
	err := db.MysqlClient.Table(r.TableName()).Where("role_id = ?", id).Find(&list).Error
	return list, err
}
