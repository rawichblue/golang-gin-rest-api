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

type GoogleAuthRequest struct {
	RedirectURL string `form:"redirect_url" binding:"required"`
}

type StateRequest struct {
	Prefix      string `json:"prefix"`
	RedirectURL string `form:"redirect_url"`
}

type GoogleUserResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}
