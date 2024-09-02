package auth

import (
	authdto "app/modules/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSvc *AuthService
}

func newController(authSvcService *AuthService) *AuthController {
	return &AuthController{
		authSvc: authSvcService,
	}
}

func (ctl *AuthController) Login(c *gin.Context) {
	var loginBody authdto.LoginBody
	if err := c.ShouldBindJSON(&loginBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, token, err := ctl.authSvc.Login(c.Request.Context(), loginBody)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "user login success",
		"token":   token,
		"user":    employee,
	})
}

// func (c *AuthController) Login(ctx *gin.Context) {
// 	var loginBody authdto.LoginBody
// 	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"code":    http.StatusBadRequest,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	user, err := c.authSvc.Login(ctx, loginBody)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{
// 			"code":    http.StatusUnauthorized,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": http.StatusOK,
// 		"data": user,
// 	})
// }
