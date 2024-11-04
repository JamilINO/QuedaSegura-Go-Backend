package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GET(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"name": "Mateus",
	})
}