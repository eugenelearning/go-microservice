package main

import (
	"api/routes"
	"api/storage"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	storage.InitDB()

	e.Use(echoprometheus.NewMiddleware("echo"))
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(500))))

	e.GET("/metrics", echoprometheus.NewHandler())
	routes.InitUserRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
