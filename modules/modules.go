package modules

import (
	"app/config"
	"app/modules/auth"
	"app/modules/employee"
	"app/modules/google"
	"app/modules/product"
	"app/modules/role"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB       *bun.DB
	Product  *product.ProductModule
	Employee *employee.ProductModule
	Auth     *auth.AuthModule
	Role     *role.RoleModule
}

func Get() *Modules {
	db := config.Database()
	google := google.New()

	return &Modules{
		DB:       db,
		Product:  product.New(db),
		Employee: employee.New(db),
		Auth:     auth.New(db, google),
		Role:     role.New(db),
	}
}
