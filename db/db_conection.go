package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Postgres = StartDB()

func StartDB () (*pgxpool.Pool){

	db_url := os.Getenv("DB_URL")

	if db_url == "" {
		db_url = "postgres://postgres:postgres@localhost:5432/quedasegura"
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	instance, err := pgxpool.New(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to postgresect to database: %v\n", err)
		os.Exit(1)
	}

	return instance
}