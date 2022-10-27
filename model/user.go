package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
)

type User struct {
	common.DBBase

	Account  string `json:"account"`  //用户名
	Password string `json:"password"` //密码
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetByAccount() error {
	return db.MysqlClient.Table(u.TableName()).Where("account = ? ", u.Account).Find(&u).Error
}

func (u *User) Create() error {
	return db.MysqlClient.Table(u.TableName()).Create(&u).Error
}
