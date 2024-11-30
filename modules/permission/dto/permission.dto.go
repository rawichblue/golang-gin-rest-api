package permissiondto

type ReqGetPermissionList struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
}

type ReqGetPermissionByID struct {
	Id int64 `uri:"id" binding:"required"`
}

type ReqStatusPermission struct {
	IsActive bool `json:"is_active"`
}
