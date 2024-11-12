package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/queue"

	/* -- Rotas -- */

	"quedasegura.com/m/v2/routes/api/contacts"
	"quedasegura.com/m/v2/routes/api/devices"
	Queda "quedasegura.com/m/v2/routes/api/queda"
	"quedasegura.com/m/v2/routes/middleware"

	"quedasegura.com/m/v2/routes/frontend/home"
	SignIn "quedasegura.com/m/v2/routes/frontend/sign_in"
	SignUp "quedasegura.com/m/v2/routes/frontend/sign_up"
)



func main() {
    

    //defer postgres.Close(context.Background())
//
    var name string
    var mac string
    var uem string
    var cem string
	rows, err := db.Postgres.Query(context.Background(), `SELECT Users.id, mac_adress, Users.email, Contacts.email FROM Devices 
INNER JOIN Users ON Users.id = Devices.foreign_id 
INNER JOIN Contacts ON Users.id = Contacts.foreign_id
;`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
    for rows.Next(){
        rows.Scan(&name, &mac, &uem, &cem)
        fmt.Printf("\n\nName: %s\nMac: %s\nEmail1: %s\nEmail2: %s\n\n", name, mac, uem, cem)
        
    }
//

	server := gin.Default()

    server.LoadHTMLGlob("./views/*")
    server.Static("/assets", "./assets")

    server.NoRoute(func(ctx *gin.Context) {
        middleware.Error(ctx, fmt.Errorf("desculpe, página não encontrada"), "A página solicitada não pôde ser encontrada", http.StatusNotFound)
    })

    /* -- GET -- */
    
    {
        server.GET("/", home.GET)
        server.GET("/sign_in", SignIn.GET)
        server.GET("/sign_up", SignUp.GET)
    }

    /* -- POST -- */    
    {
        api_group := server.Group("/api");{
            api_group.POST("", Queda.POST)
            api_group.POST("/new_contact", contacts.POST)
            api_group.POST("/update_contact", contacts.UPDATE)
            api_group.POST("/delete_contact", contacts.DELETE)
            api_group.POST("/new_device", devices.POST)
            api_group.POST("/update_device", devices.UPDATE)
            api_group.POST("/delete_device", devices.DELETE)

            /* -- DELETE -- */
            api_group.DELETE("/delete_contact", contacts.DELETE)
            api_group.DELETE("/delete_device", devices.DELETE)
        }

        server.POST("/sign_up", SignUp.POST)
        server.POST("/sign_in", SignIn.POST)
    }

	go queue.Consume()
	server.Run(":7777")
}
