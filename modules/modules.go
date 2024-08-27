package modules

import (
	"app/config"
	"app/modules/customer"
	"app/modules/product"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB       *bun.DB
	Product  *product.ProductModule
	Customer *customer.CustomerModule
}

func Get() *Modules {

	db := config.Database()

	return &Modules{
		DB:       db,
		Product:  product.New(db),
		Customer: customer.New(db),
	}
}
