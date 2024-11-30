package models

import "github.com/uptrace/bun"

type Permission struct {
	bun.BaseModel `bun:"table:permission"`

	Id          int64  `bun:",type:serial,autoincrement,pk"`
	IsActive    bool   `bun:"is_active"`
	Name        string `bun:"name"`
	Group       string `bun:"group"`
	Description string `bun:"description"`
}
