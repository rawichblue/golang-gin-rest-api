package middleware

import (
	"app/config"
	"app/models"
	"app/modules/response"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Permission(id int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Header Authorization
		header := c.Request.Header.Get("Authorization")
		hmacSampleSecret := []byte(os.Getenv("MY_SECRET_KEY"))
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		// JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		// Check Token
		if err != nil || !token.Valid {
			var message string
			if err != nil {
				if ve, ok := err.(*jwt.ValidationError); ok {
					if ve.Errors&jwt.ValidationErrorExpired != 0 {
						message = "Token is expired"
					} else {
						message = "Token is invalid"
					}
				} else {
					message = err.Error()
				}
			} else {
				message = "Invalid token"
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": message})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["userId"]
			if !ok {
				response.InternalError(c, errors.New("user not found"))
				c.Abort()
				return
			}

			db := config.Database()

			var employee models.Employee

			err := db.NewSelect().Model(&employee).Where("id = ?", userID).Scan(c)
			if err != nil {
				response.InternalError(c, errors.New("user not found"))
				c.Abort()
				return
			}

			ex, err := db.NewSelect().TableExpr("role_permission").Where("role_id = ? AND permission_id = ?", employee.RoleId, id).Exists(c)
			if err != nil {
				response.InternalError(c, errors.New("role or permission not found"))
				c.Abort()
				return
			}

			if ex {
				c.Next()
			} else {
				response.Unauthorized(c, "Unauthorized")
				c.Abort()
				return
			}

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Invalid token"})
			return
		}
	}
}
