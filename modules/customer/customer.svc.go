package customer

import (
	"app/models"
	customerdto "app/modules/customer/dto"
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type CustomerService struct {
	db *bun.DB
}

func newService(db *bun.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (s *CustomerService) Create(ctx context.Context, req customerdto.CreateCustomerdtoRequest) (*models.Customer, error) {
	m := models.Customer{
		Name:      req.Name,
		Phone:     int64(req.Phone),
		Province:  req.Province,
		CreatedAt: time.Now(),
	}
	_, err := s.db.NewInsert().Model(&m).Exec(ctx)

	return &m, err
}

func (s *CustomerService) Update(ctx context.Context, id customerdto.GetCustomerdtoByIDRequest, req customerdto.UpdateCustomerdtoRequest) (*models.Customer, error) {

	ex, err := s.db.NewSelect().Model((*models.Customer)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("customer not found")
	}

	m := models.Customer{
		ID:        id.ID,
		Name:      req.Name,
		Phone:     int64(req.Phone),
		Province:  req.Province,
		UpdatedAt: time.Now(),
	}
	_, err = s.db.NewUpdate().Model(&m).
		Set("name = ?name").
		Set("Phone = ?Phone").
		Set("Province = ?Province").
		Set("updated_at = ?updated_at").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	return &m, err
}

func (s *CustomerService) Delete(ctx context.Context, id customerdto.GetCustomerdtoByIDRequest) (*models.Customer, error) {

	exists, err := s.db.NewSelect().Model((*models.Customer)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("customer not found")
	}

	now := time.Now()

	m := models.Customer{
		ID:        id.ID,
		DeletedBy: nil,
		DeletedAt: &now,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("deleted_by = ?deleted_by").
		Set("deleted_at = ?deleted_at").
		Where("id = ?id").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *CustomerService) GetById(ctx context.Context, id customerdto.GetCustomerdtoByIDRequest) (*models.Customer, error) {
	m := models.Customer{}
	err := s.db.NewSelect().Model(&m).Where("id = ?", id.ID).Scan(ctx)
	return &m, err
}

func (s *CustomerService) GetList(ctx context.Context) ([]models.Customer, error) {
	m := []models.Customer{}
	err := s.db.NewSelect().Model(&m).Scan(ctx)
	return m, err
}
