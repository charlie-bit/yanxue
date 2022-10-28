package render

import (
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
)

type RelationResp struct {
	common.ResponseBase

	Data RelationDetail `json:"data"`
}

type RelationDetail struct {
	form.RelationReq
	Account  string `json:"account"`
	RoleName string `json:"role_name"` // 角色名称
	Alias    string `json:"alias"`
}

type RelationListResp struct {
	common.ResponseBase

	Data []RelationDetail `json:"data"`
}
