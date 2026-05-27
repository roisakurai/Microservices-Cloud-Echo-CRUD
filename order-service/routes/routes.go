package routes

import (
	"order-service/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	e.POST("/orders", handlers.CreateOrder)
	e.GET("/orders", handlers.GetAllOrders)

	e.GET("/orders/update-status", handlers.UpdateOrderStatusCronHandler)

	e.GET("/orders/:id", handlers.GetOrderByID)
	e.PUT("/orders/:id", handlers.UpdateOrder)
	e.DELETE("/orders/:id", handlers.DeleteOrder)

}
