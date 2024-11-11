package home

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/middleware"
)

type Contact struct {
	Id string
	Email string 
}

func GET(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, user, date := middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.HTML(http.StatusOK, "index.html", gin.H{

		})
		return
	}


	var id string
	var email string

	rows, err := db.Postgres.Query(context.Background(), `
	SELECT Contacts.id, Contacts.email FROM Users
	INNER JOIN Contacts ON Users.id = Contacts.foreign_id
	WHERE Users.id = $1;
	`, user)

	var contact_arr []Contact
	for rows.Next() {
		rows.Scan(&id, &email)
		contact_arr = append(contact_arr, Contact{
			Id: id,
			Email: email,
		})

	}

	if err != nil {
		middleware.Error(ctx, err, "Erro do db", http.StatusInternalServerError)
		return
	}

	

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"user": user,
		"contacts": contact_arr,
		"date": date,
	})
}