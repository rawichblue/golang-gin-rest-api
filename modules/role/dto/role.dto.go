package roledto

type ReqCreateRole struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActived   bool   `json:"is_actived"`
}
type ReqUpdateRole struct {
	ReqCreateRole
}
