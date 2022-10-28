package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
)

type ResourceColumns struct {
	common.DBBase

	form.ResourceReq
}

func (ResourceColumns) TableName() string {
	return "resource_columns"
}

func (r *ResourceColumns) Create() error {
	return db.MysqlClient.Table(r.TableName()).Create(&r).Error
}

func (r *ResourceColumns) GetByName(name string) error {
	return db.MysqlClient.Table(r.TableName()).Where("name = ?", name).Find(&r).Error
}

func (r *ResourceColumns) GetByID(id uint) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Find(&r).Error
}

func (r *ResourceColumns) Update() error {
	return db.MysqlClient.Table(r.TableName()).Save(&r).Error
}

func (r *ResourceColumns) Delete(id uint) error {
	return db.MysqlClient.Table(r.TableName()).Where("id = ?", id).Delete(&r).Error
}

func (r ResourceColumns) List() ([]common.TreeNode, error) {
	var list []*ResourceColumns
	err := db.MysqlClient.Table(r.TableName()).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return common.GenerateTree(GetResourceSlice(list)), nil
}

func (r ResourceColumns) ListByIDs(ids ...uint) ([]common.TreeNode, error) {
	var list []*ResourceColumns
	err := db.MysqlClient.Table(r.TableName()).Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return common.GenerateTree(GetResourceSlice(list)), nil
}

func GetResourceSlice(data []*ResourceColumns) (iTrees []common.ITreeNode) {
	for _, v := range data {
		iTrees = append(iTrees, v)
	}
	return iTrees
}

func (r *ResourceColumns) GetPrimKey() int {
	return int(r.ID)
}

func (r *ResourceColumns) GetParentPrimKey() int {
	return int(r.ParentID)
}

func (r *ResourceColumns) GetName() string {
	return r.Name
}

func (r *ResourceColumns) GetData() interface{} {
	return r
}

func (r *ResourceColumns) Root() bool {
	return r.ParentID == 0
}
