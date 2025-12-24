package handlers

import (
	"api/analytics"
	"api/storage"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetZStats(c echo.Context) error {
	min := -10000
	max := 10000

	randomNumber := float64(min + rand.IntN(max-min))

	storage.ZS.Add(randomNumber)

	zscore, anomaly := storage.ZS.IsAnomaly(randomNumber * 2)

	fmt.Printf("Score: %f anomaly %t\n", zscore, anomaly)

	if anomaly {
		analytics.Anomalys.Inc()
	}

	return c.JSON(http.StatusOK, anomaly)
}
