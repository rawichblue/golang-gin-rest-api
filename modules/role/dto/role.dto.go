package roledto

type ReqCreateRole struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"is_active"`
}

type RespRoleLsit struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description" `
	IsActive    bool   `json:"is_active"`
}

type ReqUpdateRole struct {
	ReqCreateRole
}

type ReqSetPermission struct {
	RoleId        int64   `json:"role_id"`
	PermissionIds []int64 `json:"permission_ids"`
}

//	type ReqPermissionId struct {
//		Id int64 `uri:"id"`
//	}
type ReqPermissionId struct {
	Id int64 `uri:"id"`
}

type ReqGetRoleList struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
}

type ReqRoleId struct {
	ID int64 `uri:"id" binding:"required"`
}

type ReqChangeStatus struct {
	IsActive bool `json:"is_active"`
}
