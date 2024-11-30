package role

import (
	"app/models"
	roledto "app/modules/role/dto"
	"context"
	"errors"

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

	_, err := s.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *RoleService) SetPermission(ctx context.Context, req roledto.ReqSetPermission) error {

	for _, per := range req.PermissionIds {
		ex, err := s.db.NewSelect().TableExpr("permission").Where("id = ? AND is_active = ?", per, true).Exists(ctx)
		if err != nil {
			return err
		}

		if !ex {
			return errors.New("status not")
		}

	}

	_, err := s.db.NewDelete().TableExpr("role_permission").Where("role_id = ?", req.RoleId).Exec(ctx)

	if err != nil {
		return err
	}

	for _, per := range req.PermissionIds {
		rolePermission := models.RolePermission{
			RoleId:       req.RoleId,
			PermissionId: per,
		}

		_, err := s.db.NewInsert().Model(&rolePermission).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RoleService) GetPermission(ctx context.Context, req roledto.ReqPermissionId) ([]int, error) {
	// []models.RolePermission
	// m := []models.RolePermission{}
	// err := s.db.NewSelect().Model(&m).Where("role_id = ?", req.Id).Scan(ctx)

	// return m, err

	var data []int

	err := s.db.NewSelect().TableExpr("role_permission").ColumnExpr("permission_id").Where("role_id = ?", req.Id).Scan(ctx, &data)

	return data, err
}

func (s *RoleService) DeleteRole(ctx context.Context, req roledto.ReqPermissionId) error {

	ex, err := s.db.NewSelect().TableExpr("employees").Where("role_id = ?", req.Id).Exists(ctx)
	if err != nil {
		return err
	}

	if ex {
		return errors.New("have user used")
	}

	_, err = s.db.NewDelete().TableExpr("role_permission").Where("role_id = ?", req.Id).Exec(ctx)

	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().TableExpr("role").Where("id = ?", req.Id).Exec(ctx)

	return err
}
