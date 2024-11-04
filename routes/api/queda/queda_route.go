package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/queue"
)

type QuedaJsonPayload struct {
	MacAddr     string `form:"mac_addr" json:"mac_addr" xml:"mac_addr"  binding:"required"`
	Date		uint32 `form:"date" json:"date" xml:"date"  binding:"required"`
	Itensity    float32 `form:"itensity" json:"itensity" xml:"itensity" binding:"required"`
}

func POST(ctx *gin.Context) {
	var json QuedaJsonPayload
	if err := ctx.ShouldBindJSON(&json); err != nil {
		print("error")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		print(json.MacAddr)
		print(json.Itensity)
	}

	fmt.Printf("Mac: %s, Itensity", json.MacAddr, )

	queue.Send(json.MacAddr, json.Date, json.Itensity)

	ctx.JSON(http.StatusOK, gin.H{
		"lorem": "ipsum",
	})
}