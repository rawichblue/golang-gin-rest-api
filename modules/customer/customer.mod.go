package customer

import "github.com/uptrace/bun"

type CustomerModule struct {
	Ctl *CustomerController
	Svc *CustomerService
}

func New(db *bun.DB) *CustomerModule {
	svc := newService(db)
	return &CustomerModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
