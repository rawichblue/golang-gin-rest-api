package role

import (
	"app/models"
	"app/modules/response"
	roledto "app/modules/role/dto"
	"context"
	"errors"
	"fmt"

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
		IsActive:    req.IsActive,
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

func (s *RoleService) GetList(ctx context.Context, req roledto.ReqGetRoleList) ([]roledto.RespRoleLsit, *response.Paginate, error) {

	var resp []models.Role

	var offset int
	if req.Page > 1 {
		offset = (req.Page - 1) * req.Size
	}

	query := s.db.NewSelect().Model(&resp).
		Column("id", "name", "description", "is_active").
		Order("created_at ASC").
		Limit(req.Size).
		Offset(offset)

	if req.Search != "" {
		search := fmt.Sprintf("%%%s%%", req.Search)
		query.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	paginate := response.Paginate{
		Page:  int64(req.Page),
		Size:  int64(req.Size),
		Total: int64(count),
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, nil, err
	}

	var result []roledto.RespRoleLsit
	for _, role := range resp {
		result = append(result, roledto.RespRoleLsit{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			IsActive:    role.IsActive,
		})
	}

	return result, &paginate, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, id roledto.ReqRoleId, req roledto.ReqChangeStatus) (*models.Role, error) {
	ex, err := s.db.NewSelect().Model((*models.Role)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("role not found")
	}

	m := models.Role{
		ID:       id.ID,
		IsActive: req.IsActive,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("is_active = ?is_active").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	return &m, err
}
