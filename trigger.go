package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/queue"
	"quedasegura.com/m/v2/routes/api/queda"
	"quedasegura.com/m/v2/routes/frontend/sign_in"
)



func main() {
    
    postgres := db.StartDB()

    defer postgres.Close(context.Background())

    var name string
    var mac string
    var uem string
    var cem string
	rows, err := postgres.Query(context.Background(), `SELECT Users.id, mac_adress, Users.email, Contacts.email FROM Devices 
INNER JOIN Users ON Users.id = Devices.foreign_id 
INNER JOIN Contacts ON Users.id = Contacts.foreign_id
WHERE mac_adress='9e:9d:17:56:60:b0';`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
    for rows.Next(){
        rows.Scan(&name, &mac, &uem, &cem)
        fmt.Printf("\n\nName: %s\nMac: %s\nEmail1: %s\nEmail2: %s\n\n", name, mac, uem, cem)
        
    }

	server := gin.Default()
    
	server.GET("/sign_in", frontend.SignIn)

    api_group := server.Group("/api");{
        api_group.POST("", api.Queda)
    }

	go queue.Consume()
	server.Run(":7777")
}
