package role

import "github.com/uptrace/bun"

type RoleModule struct {
	Ctl *RoleController
	Svc *RoleService
}

func New(db *bun.DB) *RoleModule {
	svc := newService(db)
	return &RoleModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
