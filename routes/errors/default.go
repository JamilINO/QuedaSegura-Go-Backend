package middleware_error

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, err error, msg string, status int){

	ctx.JSON(status, gin.H{
		"err": msg,
		"status": status,
		"description": err.Error(),
	})
}