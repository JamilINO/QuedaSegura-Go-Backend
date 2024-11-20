package logout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context)  {
	ctx.SetCookie("token", "", 0, "/", "", !gin.IsDebugging(), true)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

