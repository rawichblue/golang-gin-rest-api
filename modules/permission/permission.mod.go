package permission

import (
	"github.com/uptrace/bun"
)

type PermissionModule struct {
	Ctl *PermissionController
	Svc *PermissionService
}

func New(db *bun.DB) *PermissionModule {
	svc := newService(db)
	return &PermissionModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
