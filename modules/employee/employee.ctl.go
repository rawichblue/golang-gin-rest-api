package employee

import (
	employeedto "app/modules/employee/dto"
	"app/modules/response"
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

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
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

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
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

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
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

func (c *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id := employeedto.ReqGetEmployeeByID{}
	if err := ctx.BindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	data, err := c.employeeSvc.GetById(ctx, id)
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

// GetEmployeeList handles the request for listing employees with pagination and search
func (c *EmployeeController) GetEmployeeList(ctl *gin.Context) {
	var req employeedto.ReqGetEmployeeList

	if err := ctl.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctl, "Invalid request data")
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}

	employees, paginate, err := c.employeeSvc.GetList(ctl.Request.Context(), req)
	if err != nil {
		response.InternalError(ctl, err.Error())
		return
	}

	response.SuccessWithPaginate(ctl, employees, paginate)
}

// func (c *EmployeeController) GetEmployeeList(ctx *gin.Context) {
// 	data, err := c.employeeSvc.GetList(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"code":    http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"code": http.StatusOK,
// 		"data": data,
// 	})
// }
