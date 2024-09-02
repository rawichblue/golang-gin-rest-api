package routes

import (
	"app/cmd/middleware"
	"app/modules"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup, mod *modules.Modules) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", mod.Auth.Ctl.Login)
	}

	protected := r.Group("/md")
	protected.Use(middleware.CheckJwtAuth())

	employee := protected.Group("/employee")
	{
		employee.POST("/create", mod.Employee.Ctl.CreateEmployee)
		employee.PATCH("/:id", mod.Employee.Ctl.UpdateEmployee)
		employee.DELETE("/:id", mod.Employee.Ctl.DeleteEmployee)
		employee.GET("/:id", mod.Employee.Ctl.GetEmployeeById)
		employee.GET("/list", mod.Employee.Ctl.GetEmployeeList)
	}

	product := protected.Group("/product")
	{
		product.POST("/create", mod.Product.Ctl.Create)
		product.PATCH("/:id", mod.Product.Ctl.Update)
		product.DELETE("/:id", mod.Product.Ctl.Delete)
		product.GET("/:id", mod.Product.Ctl.Get)
		product.GET("/list", mod.Product.Ctl.List)
	}
}
