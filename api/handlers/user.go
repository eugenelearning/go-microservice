package handlers

import (
	"api/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := repositories.GetUserByID(c, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, user)
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	users, err := repositories.GetAllUsers(c)

	c.Logger().Print(users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, users)
	}

	return c.JSON(http.StatusOK, users)
}
