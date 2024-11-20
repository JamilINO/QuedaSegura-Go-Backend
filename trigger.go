package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"quedasegura.com/m/v2/queue"

	/* -- Rotas -- */

	"quedasegura.com/m/v2/routes/api/contacts"
	"quedasegura.com/m/v2/routes/api/devices"
	Queda "quedasegura.com/m/v2/routes/api/queda"
	"quedasegura.com/m/v2/routes/middleware"

	"quedasegura.com/m/v2/routes/frontend/home"
	SignIn "quedasegura.com/m/v2/routes/frontend/sign_in"
	SignUp "quedasegura.com/m/v2/routes/frontend/sign_up"
    "quedasegura.com/m/v2/routes/frontend/log_out"
)



func main() {
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
        server.GET("/log_out", logout.GET)
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
