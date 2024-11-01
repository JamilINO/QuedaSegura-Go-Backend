package main

import(
    "net/http"
    "github.com/gin-gonic/gin"
    "quedasegura.com/m/v2/queue"
)

func main()  {
    server := gin.Default()
    server.GET("/", func(ctx *gin.Context) {
        queue.Send()
        ctx.JSON(http.StatusOK, gin.H {
            "ok": "world",
        })
    })

    go queue.Consume()
    server.Run(":7777")
}