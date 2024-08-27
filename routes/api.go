package routes

import (
	"app/modules"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup, mod *modules.Modules) {
	product := r.Group("/product")
	{
		product.POST("/create", mod.Product.Ctl.Create)
		product.PATCH("/:id", mod.Product.Ctl.Update)
		product.DELETE("/:id", mod.Product.Ctl.Delete)
		product.GET("/:id", mod.Product.Ctl.Get)
		product.GET("/list", mod.Product.Ctl.List)
	}

	customer := r.Group("/customer")
	{
		customer.POST("/create", mod.Customer.Ctl.Create)
		customer.PATCH("/:id", mod.Customer.Ctl.Update)
		customer.DELETE("/:id", mod.Customer.Ctl.Delete)
		customer.GET("/:id", mod.Customer.Ctl.GetById)
		customer.GET("/list", mod.Customer.Ctl.GetList)
	}
}
