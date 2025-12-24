package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitCache() {

	fmt.Println(os.Getenv("REDIS_URL"))

	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to redis")
}
