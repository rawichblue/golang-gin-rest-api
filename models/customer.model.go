package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"table:customer"`

	ID        int64      `bun:",type:serial,autoincrement,pk" json:"id"`
	Name      string     `bun:"name" json:"name"`
	Phone     int64      `bun:"phone" json:"phone"`
	Province  string     `bun:"province" json:"province"`
	CreatedBy string     `bun:"created_by" json:"created_by"`
	CreatedAt time.Time  `bun:"created_at" json:"created_at"`
	DeletedBy *string    `bun:"deleted_by" json:"deleted_by,omitempty"`
	DeletedAt *time.Time `bun:"deleted_at" json:"deleted_at,omitempty"`
	UpdatedAt time.Time  `bun:"updated_at" json:"updated_at"`
}
