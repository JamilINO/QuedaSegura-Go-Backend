package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/queue"

	/* -- Rotas -- */

	Queda "quedasegura.com/m/v2/routes/api/queda"

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


    /* -- GET -- */
    
    {
        server.GET("/sign_in", SignIn.GET)
        server.GET("/sign_up", SignUp.GET)
    }

    /* -- POST -- */    
    {
        api_group := server.Group("/api");{
            api_group.POST("", Queda.POST)
        }

        server.POST("/sign_up", SignUp.POST)
    }

	go queue.Consume()
	server.Run(":7777")
}
