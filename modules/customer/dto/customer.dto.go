package customerdto

type CreateCustomerdtoRequest struct {
	Name      string  `json:"name" binding:"required"`
	Phone     float64 `json:"phone" binding:"required"`
	Province  string  `json:"province"`
	CreatedBy bool    `json:"created_by"`
	CreatedAt bool    `json:"created_at"`
}

type UpdateCustomerdtoRequest struct {
	CreateCustomerdtoRequest
}

type GetCustomerdtoByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetCustomerdtoListRequest struct {
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	SortBy   string `form:"sortBy"`
	OrderBy  string `form:"orderBy"`
	Search   string `form:"search"`
	SearchBy string `form:"searchBy"`
}
