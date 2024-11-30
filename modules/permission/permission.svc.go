package permission

import (
	"app/models"
	permissiondto "app/modules/permission/dto"
	"app/modules/response"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/uptrace/bun"
)

type PermissionService struct {
	db *bun.DB
}

func newService(db *bun.DB) *PermissionService {
	return &PermissionService{
		db: db,
	}
}

func (s *PermissionService) CreatePermission(ctx context.Context, per []models.Permission) error {
	_, err := s.db.NewInsert().Model(&per).Exec(ctx)

	return err
}

func (s *PermissionService) PermissionList(ctx context.Context, req permissiondto.ReqGetPermissionList) ([]models.Permission, *response.Paginate, error) {
	resp := []models.Permission{}

	var offset int
	if req.Page > 1 {
		offset = (req.Page - 1) * req.Size
	} else {
		offset = 0
	}

	query := s.db.NewSelect().Model(&resp)

	if req.Search != "" {
		search := fmt.Sprintf("%%%s%%", req.Search)
		query.Where("name ILIKE ? OR group ILIKE ? ", search, search)
	}

	Count, err := query.Count(ctx)

	if err != nil {
		return nil, nil, err
	}

	paginate := response.Paginate{
		Page:  int64(req.Page),
		Size:  int64(req.Size),
		Total: int64(Count),
	}

	err = query.Order("id ASC").Limit(req.Size).Offset(offset).Scan(ctx)

	log.Printf("data : %v", resp)

	if err != nil {
		return nil, nil, err
	}

	return resp, &paginate, nil
}

func (s *PermissionService) UpdatePermission(ctx context.Context, id permissiondto.ReqGetPermissionByID, req permissiondto.ReqStatusPermission) (*models.Permission, error) {
	ex, err := s.db.NewSelect().Model((*models.Permission)(nil)).Where("id = ?", id.Id).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("permission not found")
	}

	m := models.Permission{
		Id:       id.Id,
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
