package api

import (
	"errors"
	base "github.com/charlie-bit/yanxue/controller/common"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/charlie-bit/yanxue/model/form"
	"github.com/charlie-bit/yanxue/model/render"
	"github.com/gin-gonic/gin"
	"strconv"
)

var resource = &_Resource{}

type _Resource struct {
	base.SuperController
}

// New
// @Summary 创建资源
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param data body form.ResourceReq true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/resource/new [post]
func (_Resource) New(c *gin.Context) {
	var param form.ResourceReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var re model.ResourceColumns
	if err := re.GetByName(param.Name); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if re.ID != 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("资源名称重复")))
		return
	}

	re.ResourceReq = param

	if err := re.Create(); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// Update
// @Summary 更新资源
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param data body form.ResourceReq true "请示参数"
// @Param id path string true "资源 id"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/resource/{id} [put]
func (_Resource) Update(c *gin.Context) {
	id := c.Param("id")
	idi, _ := strconv.ParseUint(id, 10, 64)

	var param form.ResourceReq
	if err := c.Bind(&param); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var re model.ResourceColumns
	if err := re.GetByID(uint(idi)); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if re.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("资源不存在")))
		return
	}

	re.ResourceReq = param

	if err := re.Update(); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}

// List
// @Summary 资源列表
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Success 200 {object} render.ResourceListResp
// @Router /api/v1/resource/list [get]
func (_Resource) List(c *gin.Context) {
	var r model.ResourceColumns
	data, err := r.List()
	if err != nil {
		c.JSON(common.Failed(common.ErrSQL, err))
		return
	}

	var resp render.ResourceListResp
	resp.Data = data
	c.JSON(common.SuccessData(&resp))
}

// Del
// @Summary 删除资源
// @Schemes
// @Security Authorization
// @Description
// @accept json
// @Tags Resource
// @Param account header string true "当前账号"
// @Param id path string true "请示参数"
// @Success 200 {object} common.ResponseBase
// @Router /api/v1/resource/{id} [delete]
func (_Resource) Del(c *gin.Context) {
	id := c.Param("id")
	idi, _ := strconv.ParseUint(id, 10, 64)

	var re model.ResourceColumns
	if err := re.GetByID(uint(idi)); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	if re.ID == 0 {
		c.JSON(common.Failed(common.ErrParams, errors.New("资源不存在")))
		return
	}

	if err := re.Delete(re.ID); err != nil {
		c.JSON(common.Failed(common.ErrParams, err))
		return
	}

	var resp common.ResponseBase
	c.JSON(common.SuccessData(&resp))
}
