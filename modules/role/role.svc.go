package role

import (
	"app/models"
	roledto "app/modules/role/dto"
	"context"

	"github.com/uptrace/bun"
)

type RoleService struct {
	db *bun.DB
}

func newService(db *bun.DB) *RoleService {
	return &RoleService{
		db: db,
	}
}

func (s *RoleService) Create(ctx context.Context, req roledto.ReqCreateRole) (*models.Role, error) {
	m := models.Role{
		Name:        req.Name,
		Description: req.Description,
		IsActived:   req.IsActived,
	}

	// var err error
	_, err := s.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
