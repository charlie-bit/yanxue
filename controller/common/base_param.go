package common

import (
	"errors"
	"github.com/charlie-bit/yanxue/pkg/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SuperController struct {
	BaseParam
}

type BaseParam struct {
	Limit  int
	Offset int
}

func (SuperController) QueryLimitOffset(c *gin.Context) (dsl *BaseParam, err error) {
	dsl = &BaseParam{}

	// check limit
	var (
		reqLimit  string
		reqOffset string
		ok        bool
	)

	reqLimit, ok = c.GetQuery("limit")
	if !ok {
		dsl.Limit = constant.DefaultLimit
	} else {
		dsl.Limit, err = strconv.Atoi(reqLimit)
		if err != nil || dsl.Limit > constant.MaxLimit || dsl.Limit < constant.DefaultLimit {
			return nil, errors.New("incorrect limit value")
		}
	}

	// check offset
	reqOffset, ok = c.GetQuery("offset")
	if !ok {
		dsl.Offset = constant.DefaultOffset
	} else {
		dsl.Offset, err = strconv.Atoi(reqOffset)
		if err != nil || dsl.Offset < constant.DefaultOffset {
			return nil, errors.New("incorrect offset value")
		}
	}

	if dsl.Offset != constant.DefaultOffset {
		dsl.Offset = dsl.Offset + 1
	}

	return dsl, nil
}
