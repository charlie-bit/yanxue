package common

import (
	"github.com/charlie-bit/yanxue/log"
	"net/http"
)

type Responser interface {
	SetOK()
}

type ResponseBase struct {
	OK      bool   `json:"ok" example:"true"`
	Warning string `json:"warning,omitempty"`
	Code    Error  `json:"code" example:"0"`
	Msg     string `json:"message,omitempty"`
} // @name Response

func (rb *ResponseBase) SetOK() {
	rb.OK = true
	rb.Code = ErrNone
}

func SuccessOK() (status int, r interface{}) {
	rb := &ResponseBase{}
	rb.SetOK()
	r = rb
	status = http.StatusOK

	return
}

func SuccessData(data Responser) (status int, r interface{}) {
	if data == nil {
		return SuccessOK()
	}

	data.SetOK()
	r = data
	status = http.StatusOK

	return
}

func Failed(code Error, err error) (status int, r interface{}) {
	if err != nil {
		log.Errorf("api request failed  {code: %v, reason: %v}", code, err.Error())
	}

	msg, ok := ErrMsg[code]
	if !ok {
		msg = "unknown error"
	}

	// if it is not sensitive information, need return details to app
	if code < ErrSQL && err != nil {
		msg = err.Error()
	}
	rb := &ResponseBase{
		OK:   false,
		Code: code,
		Msg:  msg,
	}
	r = rb
	status = http.StatusOK

	return
}
