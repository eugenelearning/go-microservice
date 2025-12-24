package routes

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func InitStatsRoutes(e *echo.Echo) {
	userGroup := e.Group("/stats")

	userGroup.GET("", handlers.GetZStats)
}
