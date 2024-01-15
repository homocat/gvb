package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 // 管理员
	PermissionUser        Role = 2 // 用户
	PermissionVisitor     Role = 3 // 游客
	PermissionDisableUser Role = 4 // 被禁用户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (r Role) String() string {
	switch r {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁用户"
	default:
		return "未知角色"
	}
}
