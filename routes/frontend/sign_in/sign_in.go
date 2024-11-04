package frontend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/errors"
)


type Login struct {
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}


func GET(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "sign_in.html", gin.H{
		
	})
}

func POST(ctx *gin.Context)  {
	var user Login

	ctx.ShouldBind(&user)


	var hash []byte
	err := db.Postgres.QueryRow(context.Background(), `
	SELECT pass FROM users WHERE email = $1;
	`, user.Email).Scan(&hash)

	if err != nil {
		middleware.Error(ctx, err, "Erro do db", http.StatusInternalServerError)
	}

	ok := bcrypt.CompareHashAndPassword(hash, []byte(user.Password))

	fmt.Printf("%s", ok)

	ctx.JSON(http.StatusOK, gin.H{
		"ok": hash,
	})

}