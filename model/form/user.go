package form

import (
	"errors"
	"regexp"
)

type UserReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (user UserReq) Validate() (error, bool) {
	if err, ok := user.LoginValidate(); err != nil {
		return err, ok
	}

	if match, _ := regexp.MatchString("^\\d{6}$", user.Code); !match {
		return errors.New("验证码不正确"), false
	}

	return nil, true
}

func (user UserReq) LoginValidate() (error, bool) {
	if match, _ := regexp.MatchString("^\\d{11}$", user.Account); !match {
		return errors.New("用户名格式不正确"), false
	}

	if match, _ := regexp.MatchString("^\\d{6}$", user.Password); !match {
		return errors.New("密码格式不正确"), false
	}

	return nil, true
}

type VerifyCheckCode struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
