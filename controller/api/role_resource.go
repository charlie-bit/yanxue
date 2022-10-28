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

var uresource = &_UserResource{}

type _UserResource struct {
	base.SuperController
}

// New
// @Summary 分配用户资源权限
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param data body form.UserResourceReq true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/user_resource/new [post]
func (_UserResource) New(c *gin.Context) {
	var param form.UserResourceReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var ur model.RoleResourceColumns
	ur.RoleId = param.RoleId
	ur.ResourceId = param.ResourceId
	if err := ur.Create(); err != nil {
		c.JSON(common.Failed(common.ErrParams, errors.New("权限已存在")))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// Del
// @Summary 删除 用户资源权限
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param data body form.UserResourceReq true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/user_resource/del [delete]
func (_UserResource) Del(c *gin.Context) {
	var param form.UserResourceReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var ur model.RoleResourceColumns
	if err := ur.GetRRID(param.RoleId, param.ResourceId); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if ur.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("权限不存在")))
		return
	}

	if err := ur.Del(ur.ID); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// GetByUID
// @Summary 查看用户权限
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param role_id path string true "请示参数"
// @Success 200 {object} render.ResourceListResp
// @Router /api/v1/user_resource/{role_id} [get]
func (_UserResource) GetByUID(c *gin.Context) {
	rid := c.Param("role_id")

	var re model.RoleResourceColumns
	list, err := re.GetByRID(rid)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	resourceIds := make([]uint, 0, len(list))
	for _, v := range list {
		resourceIds = append(resourceIds, v.ResourceId)
	}

	var res model.ResourceColumns
	node, err := res.ListByIDs(resourceIds...)
	if err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp render.ResourceListResp
	resp.Data = node
	c.JSON(common.SuccessData(&resp))
}
