package api

import (
	base "github.com/charlie-bit/yanxue/controller/common"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
	"github.com/charlie-bit/yanxue/model/render"
	"github.com/gin-gonic/gin"
	"strconv"
)

var relation = &_Relation{}

type _Relation struct {
	base.SuperController
}

// Create
// @Summary 创建用户与角色关系
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Relation
// @Param account header string true "当前账号"
// @Param data body form.RelationReq true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/relation/new [post]
func (_Relation) Create(c *gin.Context) {
	var param form.RelationReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var ur model.UserRole
	if err := ur.GetByUserID(param.UserID); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	ur.UserID = param.UserID
	ur.RoleID = param.RoleID

	if ur.ID == 0 {
		if err := ur.Create(); err != nil {
			c.JSON(common.Failed(common.ErrSQL, err))
			return
		}
	} else {
		if err := ur.Update(); err != nil {
			c.JSON(common.Failed(common.ErrSQL, err))
			return
		}
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// GetUID
// @Summary 获得用户角色
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Relation
// @Param account header string true "当前账号"
// @Param user_id path string true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/relation/{user_id} [get]
func (_Relation) GetUID(c *gin.Context) {
	uid := c.Param("user_id")

	uidint, _ := strconv.ParseUint(uid, 10, 64)

	var uc model.UserRole
	if err := uc.GetByUserID(uint(uidint)); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var r model.Role
	if err := r.GetByID(strconv.Itoa(int(uc.RoleID))); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.RelationResp
	resp.Data.UserID = uc.UserID
	resp.Data.RoleID = uc.RoleID
	resp.Data.RoleName = r.RoleName
	resp.Data.Alias = r.Alias
	c.JSON(common.SuccessData(&resp))
}

// List
// @Summary 获得用户角色列表
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Relation
// @Param account header string true "当前账号"
// @Param limit query int true "页面长度"
// @Param offset query int true "记录游标"
// @Success 200 {object} render.RoleListResp
// @Router /api/v1/relation/list [get]
func (r _Relation) List(c *gin.Context) {
	param, err := r.QueryLimitOffset(c)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var rc model.UserRole
	list, err := rc.List(param.Limit, param.Offset)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.RelationListResp
	resp.Data = make([]render.RelationDetail, len(list))

	for i, userRole := range list {
		resp.Data[i].UserID = userRole.UserID
		resp.Data[i].RoleID = userRole.RoleID

		var ro model.Role
		if err = ro.GetByID(strconv.Itoa(int(userRole.RoleID))); err != nil {
			c.JSON(common.Failed(common.ErrParams, err))
			return
		}

		var u model.User
		if err = u.GetByID(userRole.UserID); err != nil {
			c.JSON(common.Failed(common.ErrParams, err))
			return
		}

		resp.Data[i].RoleName = ro.RoleName
		resp.Data[i].Alias = ro.Alias
		resp.Data[i].Account = u.Account
	}

	c.JSON(common.SuccessData(&resp))
}
