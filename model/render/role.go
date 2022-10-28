package render

import (
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
)

type RoleResp struct {
	common.ResponseBase

	Data uint `json:"data"`
}

type RoleDetailResp struct {
	common.ResponseBase

	Data model.Role `json:"data"`
}

type RoleListResp struct {
	common.ResponseBase

	Data []model.Role `json:"data"`
}
