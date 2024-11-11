package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/routes/middleware"
)

func GET(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, claims, date := middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			
		})
		return
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"user": claims,
		"date": date,
	})
}