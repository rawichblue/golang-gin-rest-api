package helper

import (
	"app/config"
	"app/models"
	"context"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

func GetUserByToken(ctx context.Context, tokenString string) (int64, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is using HMAC signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the secret key for verifying the token
		return []byte(os.Getenv("MY_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}

	// Extract the claims (including userID)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userId"]
		if !ok {
			return 0, errors.New("invalid user ID format")
		}

		db := config.Database()

		var employee models.Employee

		err := db.NewSelect().Model(&employee).Where("id = ?", userID).Scan(ctx)
		if err != nil {
			return 0, err
		}

		return employee.ID, nil
	}

	return 0, errors.New("invalid token claims")
}
