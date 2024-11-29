package employee

import (
	"app/helper"
	"app/models"
	employeedto "app/modules/employee/dto"
	"app/modules/response"
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

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

func (s *EmployeeService) Create(ctx context.Context, req employeedto.ReqCreateEmployee, image *multipart.FileHeader) (*models.Employee, error) {
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

	imagepath := ""

	if image != nil {
		url, err := helper.UploadAndResizeImage(ctx, image, "users")
		if err != nil {
			return nil, err
		}
		imagepath = url
	}

	// Hash the password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, hashErr
	}

	m := models.Employee{
		UserId:   userID,
		Name:     req.Name,
		Email:    req.Email,
		Images:   imagepath,
		Address:  req.Address,
		Phone:    req.Phone,
		Password: string(hashedPassword),
		RoleId:   req.RoleId,
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
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		RoleId:   req.RoleId,
		// Images:   req.Images,
		Address: req.Address,
		Phone:   req.Phone,
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

func (s *EmployeeService) GetById(ctx context.Context, id employeedto.ReqGetEmployeeByID) (*models.Employee, error) {

	m := models.Employee{}

	err := s.db.NewSelect().Model(&m).Where("id = ?", id.ID).Scan(ctx)

	return &m, err
}

func (s *EmployeeService) GetList(ctx context.Context, req employeedto.ReqGetEmployeeList) ([]employeedto.RespEmployee, *response.Paginate, error) {
	resp := []employeedto.RespEmployee{}

	var offset int
	if req.Page > 1 {
		offset = (req.Page - 1) * req.Size
	} else {
		offset = 0
	}

	query := s.db.NewSelect().TableExpr("employees As em").
		ColumnExpr("em.id ,em.user_id, em.email, em.name, em.images, em.address, em.phone").
		ColumnExpr("r.name As role__name, r.id As role__id").
		// ColumnExpr(`jsonb_build_object(
		// 		'id', r.id,
		// 		'name', r.name
		// 	) AS role`).
		Join("LEFT JOIN role As r On r.id = em.role_id")

	if req.Search != "" {
		search := fmt.Sprintf("%%%s%%", req.Search)
		query.Where("name ILIKE ? OR role_id ILIKE ? OR address ILIKE ?", search, search, search)
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

	err = query.Order("em.created_at ASC").Limit(req.Size).Offset(offset).Scan(ctx, &resp)

	log.Printf("data : %v", resp)

	if err != nil {
		return nil, nil, err
	}

	return resp, &paginate, nil
}

// func (s *EmployeeService) GetList(ctx context.Context, req employeedto.ReqGetEmployeeList) ([]employeedto.RespEmployee, response.Paginate, error) {
// 	var employees []models.Employee
// 	var respEmployees []employeedto.RespEmployee

// 	query := s.db.NewSelect().Model(&employees).
// 		Column("id", "user_id", "password", "name", "images", "role_id", "address", "phone").
// 		Order("created_at ASC").
// 		Limit(req.Size).
// 		Offset((req.Page - 1) * req.Size)

// 	if req.Search != "" {
// 		search := fmt.Sprintf("%%%s%%", req.Search)
// 		query.Where("name ILIKE ? OR role_id ILIKE ? OR address ILIKE ?", search, search, search)
// 	}

// 	err := query.Scan(ctx)
// 	if err != nil {
// 		return nil, response.Paginate{}, err
// 	}

// 	for _, emp := range employees {
// 		respEmp := employeedto.RespEmployee{
// 			Id:      uint(emp.ID),
// 			UserId:  emp.UserId,
// 			Name:    emp.Name,
// 			Images:  emp.Images,
// 			RoleId:  emp.RoleId,
// 			Address: emp.Address,
// 			Phone:   emp.Phone,
// 			// Password: emp.Password,
// 		}
// 		respEmployees = append(respEmployees, respEmp)
// 	}

// 	totalCount, err := s.db.NewSelect().Model((*models.Employee)(nil)).Count(ctx)
// 	if err != nil {
// 		return nil, response.Paginate{}, err
// 	}

// 	paginate := response.Paginate{
// 		Page:  int64((req.Page-1)*req.Size) + 1,
// 		Size:  int64(req.Size),
// 		Total: int64(totalCount),
// 	}

// 	return respEmployees, paginate, nil
// }

// func (s *EmployeeService) GetListsss(ctx context.Context, req employeedto.ReqGetEmployeeList) ([]employeedto.RespEmployee, response.Paginate, error) {
// 	var employees []models.Employee
// 	var respEmployees []employeedto.RespEmployee

// 	query := s.db.NewSelect().Model(&employees).
// 		Column("id", "user_id", "password", "name", "images", "role", "address", "phone").
// 		Order("created_at ASC").
// 		Limit(req.Size).
// 		Offset((req.Page - 1) * req.Size)

// 	if req.Search != "" {
// 		search := fmt.Sprintf("%%%s%%", req.Search)
// 		query.Where("name ILIKE ? OR role ILIKE ? OR address ILIKE ?", search, search, search)
// 	}

// 	err := query.Scan(ctx)
// 	if err != nil {
// 		return nil, response.Paginate{}, err
// 	}

// 	for _, emp := range employees {
// 		respEmp := employeedto.RespEmployee{
// 			Id:       uint(emp.ID),
// 			UserId:   emp.UserId,
// 			Password: emp.Password,
// 			Name:     emp.Name,
// 			Images:   emp.Images,
// 			Role:     emp.Role,
// 			Address:  emp.Address,
// 			Phone:    emp.Phone,
// 		}
// 		respEmployees = append(respEmployees, respEmp)
// 	}

// 	totalCount, err := s.db.NewSelect().Model((*models.Employee)(nil)).Count(ctx)
// 	if err != nil {
// 		return nil, Paginate{}, err
// 	}

// 	paginate := Paginate{
// 		From:  int64((pagination.CurrentPage - 1) * pagination.PerPage),
// 		Size:  int64(pagination.PerPage),
// 		Total: int64(pagination.Total),
// 	}
// 	// pagination := response.Paginate{
// 	// 	CurrentPage: req.Page,
// 	// 	PerPage:     req.Size,
// 	// 	TotalPages:  (totalCount + req.Size - 1) / req.Size,
// 	// 	Total:       totalCount,
// 	// }

// 	return respEmployees, paginate, nil
// }

// func (s *EmployeeService) GetList(ctx context.Context) ([]models.Employee, error) {

// 	m := []models.Employee{}
// 	err := s.db.NewSelect().Model(&m).Scan(ctx)

// 	return m, err
// }
