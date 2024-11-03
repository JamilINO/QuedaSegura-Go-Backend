package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	//"database/sql"
	"github.com/jackc/pgx/v5"

	"quedasegura.com/m/v2/queue"
	"quedasegura.com/m/v2/routes/api/queda"
	"quedasegura.com/m/v2/routes/frontend/sign_in"
)



func main() {
    
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/queda_segura")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	err = conn.QueryRow(context.Background(), "SELECT VERSION();").Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

    print(name)

	server := gin.Default()
    
	server.GET("/sign_in", frontend.SignIn)

    api_group := server.Group("/api");{
        api_group.POST("", api.Queda)
    }

	go queue.Consume()
	server.Run(":7777")
}
