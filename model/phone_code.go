package model

import (
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model/common"
)

type PhoneCode struct {
	common.DBBase

	Types   string `json:"types"`
	Country string `json:"country"`
	Code    string `json:"code"`
}

func (PhoneCode) TableName() string {
	return "phone_codes"
}

func (p PhoneCode) GetAll() ([]PhoneCode, error) {
	var codes []PhoneCode
	err := db.MysqlClient.Table(p.TableName()).Group("types,id").Find(&codes).Error
	if err != nil {
		return nil, err
	}
	return codes, nil
}

func (p *PhoneCode) Insert() error {
	return db.MysqlClient.Table(p.TableName()).Create(&p).Error
}
