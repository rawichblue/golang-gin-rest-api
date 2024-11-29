package auth

import (
	"app/models"
	authdto "app/modules/auth/dto"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *bun.DB
}

func newService(db *bun.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) Login(ctx context.Context, loginBody authdto.LoginBody) (*models.Employee, string, error) {
	var employee models.Employee
	err := s.db.NewSelect().
		Model(&employee).
		Where("user_id = ?", loginBody.UserId).
		Scan(ctx)

	if err != nil {
		return nil, "", err
	}

	mapData := authdto.Details{
		ID:      employee.ID,
		UserId:  employee.UserId,
		Name:    employee.Name,
		Images:  employee.Images,
		Address: employee.Address,
		Phone:   employee.Phone,
	}

	if bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginBody.Password)) != nil {
		return nil, "", errors.New("invalid credentials")
	}

	hmacSampleSecret := []byte(os.Getenv("MY_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": employee.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"data":   mapData,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, "", err
	}

	return &employee, tokenString, nil
}

func (s *AuthService) GetInfo(ctx context.Context, id int64) (*models.Employee, error) {
	var employee models.Employee
	err := s.db.NewSelect().
		Model(&employee).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *AuthService) GetUser(ctx context.Context, token string) (*authdto.GoogleUserResponse, error) {

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token)

	resBody, err := s.GetRequest(ctx, url)
	if err != nil {
		return nil, err
	}

	var googleUser authdto.GoogleUserResponse
	if err := json.Unmarshal(resBody, &googleUser); err != nil {
		return nil, err
	}

	fmt.Println("googleUser", googleUser)

	return &googleUser, nil
}

func (s *AuthService) GetRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("internal-error")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	return resBody.Bytes(), nil
}

func (s *AuthService) ExistMail(ctx context.Context, authUser *authdto.GoogleUserResponse) (string, error) {
	employee := models.Employee{}
	ex, err := s.db.NewSelect().Model(&employee).Where("email = ?", authUser.Email).Exists(ctx)
	if !ex {
		userIDNumber := 1
		if err == nil {

			var lastEmployee models.Employee
			err := s.db.NewSelect().
				Model(&lastEmployee).
				Order("id DESC").
				Limit(1).
				Scan(ctx)

			if err != nil && err.Error() != "sql: no rows in result set" {
				return "", err
			}

			fmt.Sscanf(lastEmployee.UserId, "emp-%d", &userIDNumber)
			userIDNumber++
		}
		userID := fmt.Sprintf("emp-%d", userIDNumber)

		employee = models.Employee{
			Email:    authUser.Email,
			UserId:   userID,
			Name:     authUser.Name,
			Images:   authUser.Picture,
			Address:  "",
			Phone:    0,
			Password: "",
			RoleId:   1,
		}

		_, err = s.db.NewInsert().Model(&employee).Exec(ctx)
		if err != nil {
			return "", err
		}
	}

	if err != nil {
		return "", err
	}

	err = s.db.NewSelect().
		Model(&employee).
		Where("email= ?", authUser.Email).
		Scan(ctx)

	if err != nil {
		return "", err
	}

	mapData := authdto.Details{
		ID:      employee.ID,
		UserId:  employee.UserId,
		Name:    employee.Name,
		Images:  employee.Images,
		Address: employee.Address,
		Phone:   employee.Phone,
	}

	hmacSampleSecret := []byte(os.Getenv("MY_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": employee.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"data":   mapData,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
