package api

import (
	"errors"
	base "github.com/charlie-bit/yanxue/controller/common"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
	"github.com/charlie-bit/yanxue/model/render"
	"github.com/gin-gonic/gin"
)

var role = &_Role{}

type _Role struct {
	base.SuperController
}

// New
// @Summary 创建角色
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Role
// @Param account header string true "当前账号"
// @Param data body form.Role true "请示参数"
// @Success 200 {object} render.RoleResp
// @Router /api/v1/role/new [post]
func (_Role) New(c *gin.Context) {
	var param form.Role
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var mrole model.Role
	if err := mrole.GetByName(param.RoleName); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if mrole.ID != 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("该角色已经存在")))
		return
	}

	mrole.RoleName = param.RoleName
	mrole.Alias = param.Alias
	if err := mrole.Create(); err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	var resp render.RoleResp
	resp.Data = mrole.ID
	c.JSON(common.SuccessData(&resp))
}

// GetByID
// @Summary 获得角色详情
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Role
// @Param account header string true "当前账号"
// @Param id path string true "请示参数"
// @Success 200 {object} render.RoleDetailResp
// @Router /api/v1/role/{id} [get]
func (_Role) GetByID(c *gin.Context) {
	id := c.Param("id")

	var mrole model.Role
	if err := mrole.GetByID(id); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.RoleDetailResp
	resp.Data = mrole
	c.JSON(common.SuccessData(&resp))
}

// Del
// @Summary 角色删除
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Role
// @Param account header string true "当前账号"
// @Param id path string true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/role/{id} [delete]
func (_Role) Del(c *gin.Context) {
	id := c.Param("id")

	var mrole model.Role
	if err := mrole.GetByID(id); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if mrole.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("该角色已删除")))
		return
	}

	if err := mrole.Del(id); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// List
// @Summary 获得角色列表
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Role
// @Param account header string true "当前账号"
// @Param limit query int true "页面长度"
// @Param offset query int true "记录游标"
// @Success 200 {object} render.RoleListResp
// @Router /api/v1/role/list [get]
func (r _Role) List(c *gin.Context) {
	param, err := r.QueryLimitOffset(c)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var mrole model.Role
	roles, err := mrole.List(param.Limit, param.Offset)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.RoleListResp
	resp.Data = roles

	c.JSON(common.SuccessData(&resp))
}

// Update
// @Summary 更新角色
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Role
// @Param account header string true "当前账号"
// @Param data body form.Role true "请示参数"
// @Param id path string true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/role/{id} [put]
func (_Role) Update(c *gin.Context) {
	id := c.Param("id")

	var param form.Role
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var mrole model.Role
	if err := mrole.GetByID(id); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if mrole.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("该角色不存在")))
		return
	}

	mrole.RoleName = param.RoleName
	mrole.Alias = param.Alias
	if err := mrole.Update(); err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}
