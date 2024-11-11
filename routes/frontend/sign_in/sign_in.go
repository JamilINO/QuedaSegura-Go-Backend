package frontend

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/middleware"

	"github.com/golang-jwt/jwt/v5"
)


type Login struct {
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}



func GET(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, _, _ := middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.HTML(http.StatusOK, "sign_in.html", gin.H{
		
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func POST(ctx *gin.Context)  {
	var user Login

	cookie, cookie_exits := ctx.Cookie("token")

	if cookie_exits != nil {
		fmt.Printf("Nenhum cookie encontrado!")
	} else{
		//validar cookie
		fmt.Printf("\n\n\n\n%s\n\n\n\n", cookie)

		guard, _, _ := middleware.Guard(ctx, cookie)
		
		if guard == true {
			ctx.Redirect(http.StatusPermanentRedirect, "/")
			return
		} 

		fmt.Printf("\n\n\ntk: %s\n\n\n\n", guard)
	}
	ctx.ShouldBind(&user)

	var hash []byte
	var id string
	err := db.Postgres.QueryRow(context.Background(), `
	SELECT id, pass FROM users WHERE email = $1;
	`, user.Email).Scan(&id, &hash)

	if err != nil {
		middleware.Error(ctx, err, "Erro do db", http.StatusInternalServerError)
		return
	}

	ok := bcrypt.CompareHashAndPassword(hash, []byte(user.Password))

	if ok == nil{
		tk := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
			"user": id,
			"exp": time.Now().Add(time.Hour).Unix(),
		})
	
	
		token, er := tk.SignedString([]byte(os.Getenv("TK_SECRET")))
	
	
		fmt.Printf("Cookie: %s\nToken: %s\n Err: %s", cookie, token, er)
	
		ctx.SetCookie("token", token, 3600, "/", "", !gin.IsDebugging(), true)

		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	} else{
		middleware.Error(ctx, ok, "Usuário e/ou Senha Incorreta", http.StatusUnauthorized)
		return
	}



}