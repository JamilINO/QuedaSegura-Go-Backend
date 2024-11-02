package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/queue"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main()  {
    server := gin.Default()
    server.GET("/", func(ctx *gin.Context) {
        queue.Send()
        ctx.JSON(http.StatusOK, gin.H {
            "ok": "world",
        })
    })

    server.POST("/", func(ctx *gin.Context) {

        var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
            print("error")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else{
            print(json.User)
            print(json.Password)
        }
        ctx.JSON(http.StatusOK, gin.H {
            "lorem": "ipsum",
        })
    })
 
    go queue.Consume()
    server.Run(":7777")
}