package main

import (
	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/queue"
	"quedasegura.com/m/v2/routes/api/queda"
	"quedasegura.com/m/v2/routes/frontend/index"
)



func main() {
	server := gin.Default()
    
	server.GET("/", frontend.Index)

    api_group := server.Group("/api");{
        api_group.POST("", api.Queda)
    }

	go queue.Consume()
	server.Run(":7777")
}
