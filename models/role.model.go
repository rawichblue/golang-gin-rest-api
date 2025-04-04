package models

import "github.com/uptrace/bun"

type Role struct {
	bun.BaseModel `bun:"table:role"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	IsActive    bool   `bun:"is_active"`
	CreateUpdateUnixTimestamp
	SoftDelete
}
