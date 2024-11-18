package role

import (
	"app/modules/response"
	roledto "app/modules/role/dto"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleSvc *RoleService
}

func newController(roleSvcService *RoleService) *RoleController {
	return &RoleController{
		roleSvc: roleSvcService,
	}
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	req := roledto.ReqCreateRole{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.Name == "" {
		response.BadRequest(ctx, "ใส่ชื่อมาไอเวร")
		return
	}

	data, err := c.roleSvc.Create(ctx, req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, data)
}
