package auth

import (
	"app/models"
	authdto "app/modules/auth/dto"
	"context"
	"errors"
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

	if bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginBody.Password)) != nil {
		return nil, "", errors.New("invalid credentials")
	}

	hmacSampleSecret := []byte(os.Getenv("MY_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": employee.UserId,
		"exp":    time.Now().Add(time.Minute * 1).Unix(), // Token expires in 1 minute
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, "", err
	}

	return &employee, tokenString, nil
}
