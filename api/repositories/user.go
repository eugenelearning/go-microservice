package repositories

import (
	"api/models"
	"api/storage"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) ([]models.User, error) {
	var result []models.User
	cacheKey := "users"

	data, err := storage.RDB.Get(c.Request().Context(), cacheKey).Bytes()

	if err == nil {
		err = json.Unmarshal(data, &result)

		if err == nil {
			fmt.Println("Cache hit: retrieved users from Redis")
			return result, nil
		}
	}

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

	jsonData, err := json.Marshal(result)

	storage.RDB.Set(c.Request().Context(), cacheKey, jsonData, 10*time.Minute).Err()

	if err != nil {
		log.Printf("Error setting cache data: %v", err)
	}

	return result, nil
}

func GetUserByID(c echo.Context, id string) (models.User, error) {
	item := models.User{}
	cacheKey := "users" + id

	data, cacheErr := storage.RDB.Get(c.Request().Context(), cacheKey).Bytes()

	if cacheErr == nil {
		err := json.Unmarshal(data, &item)

		if err == nil {
			fmt.Println("Cache hit: retrieved user from Redis")
			return item, nil
		}
	}

	err := storage.DB.QueryRow(c.Request().Context(), "SELECT id, name, email FROM customer WHERE id = $1", id).Scan(&item.Id, &item.Name, &item.Email)

	if err != nil {
		return item, err
	}

	jsonData, err := json.Marshal(item)

	storage.RDB.Set(c.Request().Context(), cacheKey, jsonData, 10*time.Minute).Err()

	return item, nil
}
