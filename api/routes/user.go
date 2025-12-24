package routes

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("", handlers.GetUsers)
	userGroup.GET("/:id", handlers.GetUserByID)
}
