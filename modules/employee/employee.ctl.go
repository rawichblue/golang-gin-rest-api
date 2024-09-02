package employee

import (
	employeedto "app/modules/employee/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeSvc *EmployeeService
}

func newController(employeeSvcService *EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeSvc: employeeSvcService,
	}
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	req := employeedto.ReqCreateEmployee{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.employeeSvc.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *EmployeeController) Update(ctx *gin.Context) {
	id := employeedto.ReqGetEmployeeByID{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req := employeedto.ReqUpdateEmployee{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.employeeSvc.Update(ctx, id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *EmployeeController) Delete(ctx *gin.Context) {
	id := employeedto.ReqGetEmployeeByID{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.employeeSvc.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *EmployeeController) Get(ctx *gin.Context) {
	id := employeedto.ReqGetEmployeeByID{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.employeeSvc.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func (c *EmployeeController) List(ctx *gin.Context) {

	data, err := c.employeeSvc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
