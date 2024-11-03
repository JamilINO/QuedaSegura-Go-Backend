package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"

)


func SignUp(ctx *gin.Context)  {

	

	ctx.JSON(http.StatusOK, gin.H{
		"ok": "hello",
	})
}