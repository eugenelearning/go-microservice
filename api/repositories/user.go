package repositories

import (
	"api/models"
	"api/storage"
	"log"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) ([]models.User, error) {
	var result []models.User

	rows, err := storage.DB.Query(c.Request().Context(), "SELECT id, name, email FROM customer")

	if err != nil {
		return result, nil
	}

	defer rows.Close()

	for rows.Next() {
		item := models.User{}

		if err := rows.Scan(&item.Id, &item.Name, &item.Email); err != nil {
			log.Fatalf("Error scanning row: %v\n", err)
		}
		result = append(result, item)
	}

	return result, nil
}

func GetUserByID(c echo.Context, id string) (models.User, error) {
	item := models.User{}

	err := storage.DB.QueryRow(c.Request().Context(), "SELECT id, name, email FROM customer WHERE id = $1", id).Scan(&item.Id, &item.Name, &item.Email)

	if err != nil {
		return item, err
	}

	return item, nil
}

func CreateUser(c echo.Context, payload *models.UserDTO) (models.User, error) {
	item := models.User{}

	err := storage.DB.QueryRow(c.Request().Context(), "INSERT INTO customer (name, email) VALUES ($1, $2) RETURNING id, name, email", payload.Name, payload.Email).Scan(&item.Id, &item.Name, &item.Email)

	if err != nil {
		return item, err
	}

	return item, nil
}

func UpdateUser(c echo.Context, id string, payload *models.UserDTO) (models.User, error) {
	item := models.User{}

	err := storage.DB.QueryRow(c.Request().Context(), "UPDATE customer SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email", payload.Name, payload.Email, id).Scan(&item.Id, &item.Name, &item.Email)

	if err != nil {
		return item, err
	}

	return item, nil
}

func DeleteUser(c echo.Context, id string) (string, error) {
	var result string

	err := storage.DB.QueryRow(c.Request().Context(), "DELETE FROM customer WHERE id = $1 RETURNING id", id).Scan(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
