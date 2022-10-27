package render

import "github.com/charlie-bit/yanxue/model/common"

type PhoneCodeResp struct {
	common.ResponseBase
	Data []PhoneCodeData `json:"data"`
} //@name PhoneCodeResp

type PhoneCodeData struct {
	Country string `json:"country"`
	Code    string `json:"code"`
} //@name PhoneCodeData
