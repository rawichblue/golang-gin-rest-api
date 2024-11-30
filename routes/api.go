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
		employee.POST("/create", middleware.Permission(1), mod.Employee.Ctl.CreateEmployee)
		employee.PATCH("/:id", middleware.Permission(2), mod.Employee.Ctl.UpdateEmployee)
		employee.DELETE("/:id", middleware.Permission(3), mod.Employee.Ctl.DeleteEmployee)
		employee.GET("/:id", mod.Employee.Ctl.GetEmployeeById)
		employee.GET("/list", mod.Employee.Ctl.GetEmployeeList)
	}

	role := protected.Group("/role")
	{
		role.POST("/create", mod.Role.Ctl.CreateRole)
		role.POST("/set-permission", mod.Role.Ctl.SetPermission)
		role.GET("/get-permission/:id", mod.Role.Ctl.GetPermission)
		role.DELETE("/:id", mod.Role.Ctl.DeleteRole)
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

	permission := protected.Group("/permission")
	{
		permission.POST("/create", mod.Permission.Ctl.CreatePermission)
		permission.GET("/list", mod.Permission.Ctl.PermissionList)
		permission.PATCH("/:id", mod.Permission.Ctl.PermissionChangeStatus)
	}
}
