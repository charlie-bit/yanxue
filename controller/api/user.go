package api

import (
	"context"
	"errors"
	base "github.com/charlie-bit/yanxue/controller/common"
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
	"github.com/charlie-bit/yanxue/model/render"
	"github.com/gin-gonic/gin"
	"regexp"
	"time"
)

var user = &_User{}

type _User struct {
}

// Register
// @Summary 用户注册, 手机注册方式
// @Schemes
// @Description
// @accept json
// @Tags User
// @Param data body form.UserReq true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/register [post]
func (_User) Register(c *gin.Context) {
	var param form.UserReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if err, ok := param.Validate(); !ok {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if param.Code != db.RedisClient.Get(context.Background(), "sms_"+param.Account).Val() {
		c.JSON(common.Failed(common.ErrParams, nil))
		return
	}

	u := model.User{}
	u.Account = param.Account
	u.Password = base.GeneratePassword(param.Password)

	if err := u.GetByAccount(); err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	if u.ID != 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("账号已存在")))
		return
	}

	if err := u.Create(); err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// Login
// @Summary 用户登录
// @Schemes
// @Description
// @accept json
// @Tags User
// @Param data body form.UserReq true "请示参数"
// @Success 200 {object} render.UserLogin
// @Router /api/v1/login [post]
func (_User) Login(c *gin.Context) {
	var param form.UserReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if err, ok := param.LoginValidate(); !ok {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var u model.User
	u.Account = param.Account
	if err := u.GetByAccount(); err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	if u.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("账号不存在")))
		return
	}

	if err := base.VerifyPassword(param.Password, u.Password); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	token, err := base.NewJWTToken(param.Account)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if err = db.RedisClient.Set(context.Background(), param.Account, token.Token, time.Hour*1).Err(); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.UserLogin
	resp.Data = token.Token
	c.JSON(common.SuccessData(&resp))
}

// SignOut
// @Summary 用户退出登录
// @Security Authorization
// @Schemes
// @Description
// @accept json
// @Tags User
// @Param account header string true "当前账号"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/sign_out [get]
func (_User) SignOut(c *gin.Context) {
	account := c.Request.Header.Get("account")
	if err := db.RedisClient.Del(context.Background(), account).Err(); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// GetCheckCode
// @Summary 获取手机验证码
// @Schemes
// @Description
// @accept json
// @Tags User
// @Param phone path string true "手机号"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/check_code/{phone} [get]
func (_User) GetCheckCode(c *gin.Context) {
	phone := c.Param("phone")
	if match, _ := regexp.MatchString("^\\d{11}$", phone); !match {
		c.JSON(common.Failed(common.ErrParams, nil))
	}

	if err := base.SendCode(phone); err != nil {
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// VerifyCheckCode
// @Summary 验证手机验证码
// @Schemes
// @Deprecated
// @Description
// @accept json
// @Tags User
// @Param data body form.VerifyCheckCode true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/verify_check_code [post]
func (_User) VerifyCheckCode(c *gin.Context) {
	var param form.VerifyCheckCode
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if param.Code != db.RedisClient.Get(context.Background(), "sms_"+param.Phone).Val() {
		c.JSON(common.Failed(common.ErrParams, nil))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// PhoneCode
// @Summary 获得手机区号
// @Schemes
// @Description
// @accept json
// @Tags User
// @Success 200 {object} render.PhoneCodeResp
// @Router /api/v1/phone_code [get]
func (_User) PhoneCode(c *gin.Context) {
	var codes model.PhoneCode
	list, err := codes.GetAll()
	if err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	var resp render.PhoneCodeResp
	resp.Data = make([]render.PhoneCodeData, len(list))
	for i, code := range list {
		resp.Data[i].Code = code.Code
		resp.Data[i].Country = code.Country
	}
	c.JSON(common.SuccessData(&resp))
}
