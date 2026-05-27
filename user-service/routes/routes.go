package routes

import (
	"user-service/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/users", handlers.CreateUser)
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/users/:id", handlers.GetUserByID)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
}
