package form

type RelationReq struct {
	UserID uint `json:"user_id" v:"required|min:1#用户id不能为空|用户id长度有误"`           // 用户id
	RoleID uint `json:"role_id" v:"required|min-length:1#待关联角色id不能为空|关联一个角色id"` // 角色id
}
