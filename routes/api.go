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

	product := protected.Group("/product")
	{
		product.POST("/create", mod.Product.Ctl.Create)
		product.PATCH("/:id", mod.Product.Ctl.Update)
		product.DELETE("/:id", mod.Product.Ctl.Delete)
		product.GET("/:id", mod.Product.Ctl.Get)
		product.GET("/list", mod.Product.Ctl.List)
	}

	employee := protected.Group("/employee")
	{
		employee.POST("/create", mod.Employee.Ctl.Create)
		employee.PATCH("/:id", mod.Employee.Ctl.Update)
		employee.DELETE("/:id", mod.Employee.Ctl.Delete)
		employee.GET("/:id", mod.Employee.Ctl.Get)
		employee.GET("/list", mod.Employee.Ctl.List)
	}
}
