package models

import "github.com/uptrace/bun"

type RolePermission struct {
	bun.BaseModel `bun:"table:role_permission"`

	RoleId       int64 `bun:"role_id"`
	PermissionId int64 `bun:"permission_id"`
}
