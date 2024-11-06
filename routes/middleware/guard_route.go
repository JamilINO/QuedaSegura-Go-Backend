package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	User    string `json:"user"`
	Exp string `json:"exp"`
	jwt.StandardClaims
}

func Guard(ctx *gin.Context, my string) bool {
	parser, err := jwt.ParseWithClaims(my, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("TK_SECRET")), nil
	})
	if claims, ok := parser.Claims.(UserClaims); ok {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
	

	if err != nil || !parser.Valid {
		ctx.SetCookie("token", "", 0, "/", "", !gin.IsDebugging(), true)
		return false
	} else {
		//ctx.Redirect(http.StatusOK, "/")
		return true
	}
}

