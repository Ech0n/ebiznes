package routes

import (
	"project/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.GET("/categories", controllers.GetCategories)
	e.POST("/categories", controllers.CreateCategory)

	e.POST("/payments", controllers.ConfirmOrder)
}
