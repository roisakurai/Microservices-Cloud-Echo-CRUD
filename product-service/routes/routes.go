package routes

import (
	"product-service/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/products", handlers.CreateProduct)
	e.GET("/products", handlers.GetAllProducts)
	e.GET("/products/:id", handlers.GetProductByID)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)

}
