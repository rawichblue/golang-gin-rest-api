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
		auth.GET("google/login", mod.Auth.Ctl.GoogleLogin)
		auth.GET("google/callback", mod.Auth.Ctl.GoogleCallback)
		auth.GET("/getInfo", mod.Auth.Ctl.GetInfo)
		auth.POST("/login", mod.Auth.Ctl.Login)
	}

	protected := r.Group("/md")
	protected.Use(middleware.CheckJwtAuth())

	employee := r.Group("/employee")
	{
		employee.GET("/:id", mod.Employee.Ctl.GetEmployeeById)
		employee.GET("/list", mod.Employee.Ctl.GetEmployeeList)
		employee.POST("/create", middleware.Permission(1), mod.Employee.Ctl.CreateEmployee)
		employee.PATCH("/:id", middleware.Permission(2), mod.Employee.Ctl.UpdateEmployee)
		employee.DELETE("/:id", middleware.Permission(3), mod.Employee.Ctl.DeleteEmployee)
	}

	role := protected.Group("/role")
	{
		role.GET("/get-permission/:id", mod.Role.Ctl.GetPermission)
		role.GET("/list", mod.Role.Ctl.GetRoleList)
		role.POST("/create", mod.Role.Ctl.CreateRole)
		role.POST("/set-permission", mod.Role.Ctl.SetPermission)
		role.PATCH("toggle-status/:id", mod.Role.Ctl.RoleChangeStatus)
		role.PATCH("/:id", mod.Role.Ctl.Update)
		role.DELETE("/:id", mod.Role.Ctl.DeleteRole)
	}

	product := protected.Group("/product")
	{
		product.GET("/:id", mod.Product.Ctl.Get)
		product.GET("/list", mod.Product.Ctl.List)
		product.POST("/create", mod.Product.Ctl.Create)
		product.PATCH("/:id", mod.Product.Ctl.Update)
		product.DELETE("/:id", mod.Product.Ctl.Delete)
	}

	permission := protected.Group("/permission")
	{
		permission.GET("/list", mod.Permission.Ctl.PermissionList)
		permission.POST("/create", mod.Permission.Ctl.CreatePermission)
		permission.PATCH("/:id", mod.Permission.Ctl.PermissionChangeStatus)
	}
}
