package role

import (
	"app/modules/response"
	roledto "app/modules/role/dto"
	"net/http"

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

func (c *RoleController) SetPermission(ctx *gin.Context) {
	req := roledto.ReqSetPermission{}
	if err := ctx.Bind(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.RoleId == 0 {
		response.BadRequest(ctx, "ใส่เลขมาด้วย จ้า")
		return
	}

	err := c.roleSvc.SetPermission(ctx, req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *RoleController) GetPermission(ctx *gin.Context) {
	req := roledto.ReqPermissionId{}
	if err := ctx.BindUri(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.Id == 0 {
		response.BadRequest(ctx, "ใส่เลขมาด้วย จ้า")
		return
	}

	data, err := c.roleSvc.GetPermission(ctx, req)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.Success(ctx, data)
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	id := roledto.ReqPermissionId{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := c.roleSvc.DeleteRole(ctx, id)
	if err != nil {
		response.InternalError(ctx, err.Error())
		return
	}

	response.Success(ctx, nil)
}
