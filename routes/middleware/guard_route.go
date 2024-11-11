package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct{
	User string `json:"user" binding:"required"`;
	Exp int `json:"exp" binding:"required"`;
	jwt.Claims
}

func Guard(ctx *gin.Context, str_token string) (bool, string, int32) {

	if str_token == "" {
		return false, "", 0
	}

	parser, err := jwt.Parse(str_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("str_token_secret_key")
		return []byte(os.Getenv("TK_SECRET")), nil
	})

	claims, _ := parser.Claims.(jwt.MapClaims);
	fmt.Printf("Err: \n\nClaim: %s", claims)
	expiration := time.Unix(int64(claims["exp"].(float64)), 0)

	if err != nil || !parser.Valid || !time.Now().Before(expiration) {
		ctx.SetCookie("token", "", 0, "/", "", !gin.IsDebugging(), true)
		return false, "",0
	} else {
		//ctx.Redirect(http.StatusOK, "/")
		return true, claims["user"].(string), int32(claims["exp"].(float64))
	}
}

