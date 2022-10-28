package form

type ResourceReq struct {
	ParentID int64  `json:"parent_id"`                                                      // 父资源id
	Name     string `json:"name" v:"name@required#资源名称不能为空"`                                // 资源名称
	Alias    string `json:"alias"`                                                          // 资源别称
	Url      string `json:"url"`                                                            // 资源路径
	Enable   bool   `json:"enable"`                                                         // 0:不显示，1显示
	Icon     string `json:"icon"`                                                           // 资源图标
	Type     string `json:"type" v:"type@required|in:menu,button,link#资源类型不能为空|请选择正确的资源类型"` // 资源类型：menu-菜单；button-按钮；link-链接
	Sn       int    `json:"sn"`                                                             // 排序
}

type UserResourceReq struct {
	RoleId     uint `json:"role_id"`     // 角色id
	ResourceId uint `json:"resource_id"` // 资源id
}
