package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GET(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "sign_in.html", gin.H{
		"name": "Mateus",
	})
}