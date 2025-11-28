package routes

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.GET("", handlers.GetUsers)
	userGroup.GET("/:id", handlers.GetUserByID)
	userGroup.PUT("/:id", handlers.UpdateUserByID)
	userGroup.DELETE("/:id", handlers.DeleteUserByID)
	userGroup.POST("", handlers.CreateUser)
}
