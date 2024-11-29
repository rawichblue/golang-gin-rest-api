package authdto

type LoginBody struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Details struct {
	ID      int64  `json:"id"`
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	Images  string `json:"images"`
	Address string `json:"address"`
	Phone   int64  `json:"phone"`
}
