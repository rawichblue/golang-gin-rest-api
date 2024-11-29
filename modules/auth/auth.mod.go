package auth

import (
	"app/modules/google"

	"github.com/uptrace/bun"
)

type AuthModule struct {
	Ctl *AuthController
	Svc *AuthService
}

func New(db *bun.DB,
	google *google.GoogleModule) *AuthModule {
	svc := newService(db)
	return &AuthModule{
		Ctl: newController(svc, google),
		Svc: svc,
	}
}
