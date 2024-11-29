package auth

import (
	"app/helper"
	authdto "app/modules/auth/dto"
	"app/modules/google"
	"app/modules/response"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSvc *AuthService
	google  *google.GoogleModule
}

func newController(authSvcService *AuthService, google *google.GoogleModule) *AuthController {
	return &AuthController{
		authSvc: authSvcService,
		google:  google,
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

func (ctl *AuthController) GoogleLogin(c *gin.Context) {
	req := authdto.GoogleAuthRequest{}
	if err := c.Bind(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	state := encodeState("login", req.RedirectURL)
	url := ctl.google.Svc.Oauth().AuthCodeURL(state)

	c.Redirect(http.StatusFound, url)
}

func encodeState(prefix string, redirectURL string) string {

	state := authdto.StateRequest{
		Prefix:      prefix,
		RedirectURL: redirectURL,
	}

	stateJSON, _ := json.Marshal(state)

	return base64.URLEncoding.EncodeToString(stateJSON)
}

func (ctl *AuthController) GoogleCallback(ctx *gin.Context) {
	code, stateReq, err := extractCodeAndState(ctx)
	if err != nil {
		response.BadRequest(ctx, err)
		return

	}

	state, err := decodeState(stateReq)
	if err != nil {
		response.InternalError(ctx, err)
		return
	}

	tokenAuth, err := ctl.google.Svc.Oauth().Exchange(ctx, code)
	if err != nil {
		response.InternalError(ctx, err)
		return
	}

	authUser, err := ctl.authSvc.GetUser(ctx, tokenAuth.AccessToken)
	if err != nil {
		response.InternalError(ctx, err)
		return
	}

	token, err := ctl.authSvc.ExistMail(ctx, authUser)
	if err != nil {
		response.InternalError(ctx, err)
		return
	}

	// c.setTokenCookie(ctx, token, state.RedirectURL)
	ctx.Redirect(http.StatusFound, state.RedirectURL+"?token="+token)
}

func extractCodeAndState(ctx *gin.Context) (string, string, error) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	if code == "" || state == "" {
		return "", "", errors.New("unauthorized")
	}
	return code, state, nil
}

func decodeState(stateReq string) (*authdto.StateRequest, error) {
	stateBytes, err := base64.URLEncoding.DecodeString(stateReq)
	if err != nil {
		return nil, err
	}
	var stateRequest authdto.StateRequest
	if err := json.Unmarshal(stateBytes, &stateRequest); err != nil {
		return nil, err
	}
	return &stateRequest, nil
}
