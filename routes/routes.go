package routes

import (
	"task-intern-product-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	routes := gin.Default()

	routes.Use(cors.Default())
	routes.Static("/uploads", "./uploads")
	
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