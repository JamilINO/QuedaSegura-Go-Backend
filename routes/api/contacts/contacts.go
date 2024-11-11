package contacts

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/middleware"
)

type NewContact struct {
	TargetEmail string `form:"email" json:"email" xml:"email"  binding:"required"`
}

func POST(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, user, _:= middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.Redirect(http.StatusUnauthorized,"/sign_in")
		return
	}

	var contact NewContact

	ctx.ShouldBind(&contact)


	db.Postgres.QueryRow(context.Background(), `
	INSERT INTO Contacts (id, foreign_id, email)
	VALUES (
		gen_random_uuid(),
		$1,
		$2 
	)
	`, user, contact.TargetEmail)

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

type ContactId struct {
	TargetId string `form:"contact_id" json:"contact_id" xml:"contact_id"  binding:"required"`
}

func DELETE(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, user, _:= middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.Redirect(http.StatusUnauthorized,"/sign_in")
		return
	}

	var contact_id ContactId

	ctx.ShouldBind(&contact_id)


	db.Postgres.QueryRow(context.Background(), `
	INSERT INTO Contacts (id, foreign_id, email)
	VALUES (
		gen_random_uuid(),
		$1,
		$2 
	)
	`, user, contact_id.TargetId)

	ctx.Redirect(http.StatusMovedPermanently, "/")
}