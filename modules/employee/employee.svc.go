package employee

import (
	"app/models"
	employeedto "app/modules/employee/dto"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/uptrace/bun"
)

type EmployeeService struct {
	db *bun.DB
}

func newService(db *bun.DB) *EmployeeService {
	return &EmployeeService{
		db: db,
	}
}

func (s *EmployeeService) Create(ctx context.Context, req employeedto.ReqCreateEmployee) (*models.Employee, error) {
	var lastEmployee models.Employee
	err := s.db.NewSelect().
		Model(&lastEmployee).
		Order("id DESC").
		Limit(1).
		Scan(ctx)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	userIDNumber := 1
	if err == nil {
		fmt.Sscanf(lastEmployee.UserId, "emp-%d", &userIDNumber)
		userIDNumber++
	}
	userID := fmt.Sprintf("emp-%d", userIDNumber)

	// Hash the password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, hashErr
	}

	m := models.Employee{
		UserId:   userID,
		Name:     req.Name,
		Images:   req.Images,
		Address:  req.Address,
		Phone:    req.Phone,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	_, err = s.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *EmployeeService) Update(ctx context.Context, id employeedto.ReqGetEmployeeByID, req employeedto.ReqUpdateEmployee) (*models.Employee, error) {
	ex, err := s.db.NewSelect().Model((*models.Employee)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("employee not found")
	}

	m := models.Employee{
		ID:       id.ID,
		UserId:   req.UserId,
		Password: req.Password,
		Name:     req.Name,
		Role:     req.Role,
		Images:   req.Images,
		Address:  req.Address,
		Phone:    req.Phone,
	}

	_, err = s.db.NewUpdate().Model(&m).
		Set("name = ?name").
		Set("password = ?password").
		Set("role = ?role").
		Set("images = ?images").
		Set("address = ?address").
		Set("phone = ?phone").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	return &m, err
}

func (s *EmployeeService) Delete(ctx context.Context, id employeedto.ReqGetEmployeeByID) (*models.Employee, error) {

	ex, err := s.db.NewSelect().Model((*models.Employee)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("employee not found")
	}

	_, err = s.db.NewDelete().Model((*models.Employee)(nil)).Where("id = ?", id.ID).Exec(ctx)
	return nil, err
}

func (s *EmployeeService) Get(ctx context.Context, id employeedto.ReqGetEmployeeByID) (*models.Employee, error) {

	m := models.Employee{}

	err := s.db.NewSelect().Model(&m).Where("id = ?", id.ID).Scan(ctx)

	return &m, err
}

func (s *EmployeeService) List(ctx context.Context) ([]models.Employee, error) {

	m := []models.Employee{}
	err := s.db.NewSelect().Model(&m).Scan(ctx)

	return m, err
}
