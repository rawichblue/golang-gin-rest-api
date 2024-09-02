package authdto

type LoginBody struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}
