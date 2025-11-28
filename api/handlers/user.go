package handlers

import (
	"api/models"
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

func CreateUser(c echo.Context) error {
	payload := new(models.UserDTO)

	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user, err := repositories.CreateUser(c, payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, user)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c echo.Context) error {
	payload := new(models.UserDTO)
	id := c.Param("id")

	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if _, err := repositories.GetUserByID(c, id); err != nil {
		return c.JSON(http.StatusNotFound, id)
	}

	user, err := repositories.UpdateUser(c, id, payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"id": id})
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	deleted, err := repositories.DeleteUser(c, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, deleted)
	}

	return c.JSON(http.StatusOK, map[string]string{"id": deleted})
}
