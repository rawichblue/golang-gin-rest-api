package routes

import (
	"app/cmd/middleware"
	"app/modules"

	// "app/routes"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup, mod *modules.Modules) {

	auth := r.Group("/auth")
	{
		auth.GET("/getInfo", mod.Auth.Ctl.GetInfo)
		auth.POST("/login", mod.Auth.Ctl.Login)
		auth.GET("google/login", mod.Auth.Ctl.GoogleLogin)
		auth.GET("google/callback", mod.Auth.Ctl.GoogleCallback)

	}

	protected := r.Group("/md")
	protected.Use(middleware.CheckJwtAuth())

	employee := r.Group("/employee")
	{
		employee.POST("/create", mod.Employee.Ctl.CreateEmployee)
		employee.PATCH("/:id", mod.Employee.Ctl.UpdateEmployee)
		employee.DELETE("/:id", mod.Employee.Ctl.DeleteEmployee)
		employee.GET("/:id", mod.Employee.Ctl.GetEmployeeById)
		employee.GET("/list", mod.Employee.Ctl.GetEmployeeList)
	}

	role := protected.Group("/role")
	{
		role.POST("/create", mod.Role.Ctl.CreateRole)
		// role.PATCH("/:id", mod.Role.Ctl.UpdateRole)
		// role.DELETE("/:id", mod.Role.Ctl.DeleteRole)
		// role.GET("/:id", mod.Role.Ctl.GetRoleById)
		// role.GET("/list", mod.Role.Ctl.GetRoleList)
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
