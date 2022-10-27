package render

import "github.com/charlie-bit/yanxue/model/common"

type UserLogin struct {
	common.ResponseBase

	Data string `json:"data"`
}
