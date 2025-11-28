package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err.Error())
	}

	DB = pool

	err = DB.Ping(context.Background())

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}
