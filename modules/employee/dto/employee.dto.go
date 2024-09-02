package employeedto

import "app/models"

// Request
type ReqCreateEmployee struct {
	UserId   string `json:"userId"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Images   string `json:"images"`
	Address  string `json:"address"`
	Phone    int64  `json:"phone"`
}

type ReqUpdateEmployee struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Images   string `json:"images"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Phone    int64  `json:"phone"`
	models.UpdateUnixTimestamp
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

// Response
type RespEmployee struct {
	Id      uint   `json:"id"`
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	Images  string `json:"images"`
	Role    string `json:"role"`
	Address string `json:"address"`
	Phone   int64  `json:"phone"`
	// Password string `json:"password"`
}
