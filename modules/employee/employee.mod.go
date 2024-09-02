package employee

import "github.com/uptrace/bun"

type ProductModule struct {
	Ctl *EmployeeController
	Svc *EmployeeService
}

func New(db *bun.DB) *ProductModule {
	svc := newService(db)
	return &ProductModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
