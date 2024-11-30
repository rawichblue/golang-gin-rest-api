package permission

import (
	"app/models"
	permissiondto "app/modules/permission/dto"
	"app/modules/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	permissionSvc *PermissionService
}

func newController(permissionSvcService *PermissionService) *PermissionController {
	return &PermissionController{
		permissionSvc: permissionSvcService,
	}
}

var permission = []models.Permission{
	{Id: 1, IsActive: true, Name: "employee_Create", Group: "employee", Description: "employee_Create"},
	{Id: 2, IsActive: true, Name: "employee_Update", Group: "employee", Description: "employee_Update"},
	{Id: 3, IsActive: true, Name: "employee_Delete", Group: "employee", Description: "employee_Delete"},
}

func (ctl PermissionController) CreatePermission(c *gin.Context) {
	err := ctl.permissionSvc.CreatePermission(c.Request.Context(), permission)
	if err != nil {
		response.InternalError(c, err.Error())
	}

	response.Success(c, nil)
}

func (ctl *PermissionController) PermissionList(c *gin.Context) {
	var req permissiondto.ReqGetPermissionList

	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, "Invalid request data")
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	employees, paginate, err := ctl.permissionSvc.PermissionList(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.SuccessWithPaginate(c, employees, paginate)
}

func (ctl *PermissionController) PermissionChangeStatus(c *gin.Context) {
	id := permissiondto.ReqGetPermissionByID{}

	if err := c.BindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	status := permissiondto.ReqStatusPermission{}
	if err := c.Bind(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := ctl.permissionSvc.UpdatePermission(c, id, status)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, data)
}
