package main

import (
	"api/analytics"
	"api/routes"
	"api/storage"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	storage.InitDB()
	storage.InitCache()
	storage.InitZScore(2)

	analytics.Init()

	e.Use(echoprometheus.NewMiddleware("echo"))
	e.Use(middleware.Logger())

	e.GET("/metrics", echoprometheus.NewHandler())
	routes.InitUserRoutes(e)
	routes.InitStatsRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
