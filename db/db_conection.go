package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func StartDB () (*pgx.Conn){

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	postgres, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/quedasegura")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to postgresect to database: %v\n", err)
		os.Exit(1)
	}

	return postgres
}