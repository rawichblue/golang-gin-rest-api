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
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	SortBy   string `form:"sortBy"`
	OrderBy  string `form:"orderBy"`
	Search   string `form:"search"`
	SearchBy string `form:"searchBy"`
}

// Response
