package auth

import "github.com/uptrace/bun"

type AuthModule struct {
	Ctl *AuthController
	Svc *AuthService
}

func New(db *bun.DB) *AuthModule {
	svc := newService(db)
	return &AuthModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
