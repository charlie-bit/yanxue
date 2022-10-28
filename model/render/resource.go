package render

import "github.com/charlie-bit/yanxue/model/common"

type ResourceListResp struct {
	common.ResponseBase

	Data []common.TreeNode `json:"data"`
}

// ResourceTree 资源树形结构
type ResourceTree struct {
	ResourceInfo
	Child []ResourceTree `json:"child"` // 子资源
}

// ResourceRoleTree 角色对应的资源树形结构
type ResourceRoleTree struct {
	ResourceInfo
	Checked bool               `json:"checked"` // 是否选中
	Child   []ResourceRoleTree `json:"child"`   // 子资源
}

type ResourceInfo struct {
	Id       int64  `json:"id"`        // 主键id
	ParentId int64  `json:"parent_id"` // 父资源id
	Name     string `json:"name"`      // 资源名称
	Alias    string `json:"alias"`     // 资源别称
	Url      string `json:"url"`       // 资源路径
	Enable   bool   `json:"enable"`    // 0:不显示，1显示
	Icon     string `json:"icon"`      // 资源图标
	Type     string `json:"type"`      // 资源类型：menu-菜单；button-按钮；link-链接
	Sn       int    `json:"sn"`        // 排序
}
