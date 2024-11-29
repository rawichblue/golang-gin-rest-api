package auth

import (
	"app/helper"
	authdto "app/modules/auth/dto"
	"app/modules/response"
	"net/http"
	"strings"

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

func (ctl *AuthController) GetInfo(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)

	userID, err := helper.GetUserByToken(c, tokenString)
	if err != nil {
		response.BadRequest(c, err.Error())
	}

	employee, err := ctl.authSvc.GetInfo(c.Request.Context(), userID)

	if err != nil {
		response.BadRequest(c, err.Error())
	}

	response.Success(c, employee)
}

// func (c *AuthController) Login(ctx *gin.Context) {
//  var loginBody authdto.LoginBody
//  if err := ctx.ShouldBindJSON(&loginBody); err != nil {
//      ctx.JSON(http.StatusBadRequest, gin.H{
//          "code":    http.StatusBadRequest,
//          "message": err.Error(),
//      })
//      return
//  }

//  user, err := c.authSvc.Login(ctx, loginBody)
//  if err != nil {
//      ctx.JSON(http.StatusUnauthorized, gin.H{
//          "code":    http.StatusUnauthorized,
//          "message": err.Error(),
//      })
//      return
//  }

//  ctx.JSON(http.StatusOK, gin.H{
//      "code": http.StatusOK,
//      "data": user,
//  })
// }
