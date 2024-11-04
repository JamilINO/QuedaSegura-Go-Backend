package frontend

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"quedasegura.com/m/v2/db"
)


func GET(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "sign_up.html", gin.H{
		"hello": "world",
	})
}

type User struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm"  binding:"required"`

}

func POST(ctx *gin.Context)  {

	var new_user User

	ctx.ShouldBind(&new_user)

	fmt.Printf(new_user.Email)
	fmt.Printf(new_user.Password)
	fmt.Printf(new_user.PasswordConfirm)
	fmt.Printf(new_user.Name)

	if new_user.Password != new_user.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "As senhas são diferentes",
		})
		return
	}

	var user_exists int

	err := db.Postgres.QueryRow(context.Background(), `SELECT COUNT(1) FROM users WHERE email = $1`, new_user.Email).Scan(&user_exists)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "algo deu errado erro no db",
		})
		return
	}

	if user_exists != 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "user exists",
		})
		return
	} 

	hash, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), 16)
	if err != nil {
		fmt.Printf("%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "algo deu errado erro no hash",
		})
		return
	}
	fmt.Printf("Hash: %s", string(hash))

	_, db_err := db.Postgres.Query(context.Background(), `
	INSERT INTO Users(id, email, pass, real_name)
	VALUES (
		gen_random_uuid (),
		$1,
		$2,
		$3
	);`, new_user.Email, hash, new_user.Name)

	if db_err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "algo deu errado erro no db",
		})
		return
	}

	fmt.Printf("\nThis user have email %d\n", user_exists)
	ctx.Redirect(http.StatusFound, "/sign_in")


}