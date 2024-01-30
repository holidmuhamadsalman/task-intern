package routes

import (
	"task-intern-product-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	routes := gin.Default()

	product := routes.Group("/api/v1/product")
	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/", controllers.GetProduct)
		product.GET("/:id", controllers.GetProductById)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	return routes
}