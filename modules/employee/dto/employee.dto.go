package employeedto

// Request
type ReqCreateEmployee struct {
	UserId   string `json:"userId"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	RoleId   int64  `json:"role_id" binding:"required"`
	Images   string `json:"images"`
	Address  string `json:"address"`
	Phone    int64  `json:"phone"`
}

type ReqUpdateEmployee struct {
	ReqCreateEmployee
}

type ReqGetEmployeeByID struct {
	ID int64 `uri:"id" binding:"required"`
}

type ReqGetEmployeeList struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
}

type ReqGetListEmployees struct {
	Search string `form:"search" json:"search"`
	Limit  int    `form:"limit" json:"limit" binding:"required"`
	Page   int    `form:"page" json:"page" binding:"required"`
}

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"Name"`
}

// Response
type RespEmployee struct {
	ID      int64  `json:"id"`
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	Images  string `json:"images"`
	Role    Role   `json:"role"`
	Address string `json:"address"`
	Phone   int64  `json:"phone"`
	// Password string `json:"password"`
}
