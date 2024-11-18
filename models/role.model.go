package models

import "github.com/uptrace/bun"

type Role struct {
	bun.BaseModel `bun:"table:role"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	IsActived   bool   `bun:"is_actived,default:false"`
	CreateUpdateUnixTimestamp
	SoftDelete
}
