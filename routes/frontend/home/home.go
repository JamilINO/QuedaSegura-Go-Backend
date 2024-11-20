package home

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/middleware"
)

type Contact struct {
	Id string
	Email string 
}

type Device struct {
	Id string
	MacAddr string
}

func GET(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, user, date := middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.HTML(http.StatusOK, "index.html", gin.H{

		})
		return
	}

	/* -- Inicia READ Contatos -- */
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
	/* -- Encerra READ Contatos -- */


	/* -- Inicia READ Devices -- */

	var devices_arr []Device

	var device_id string
	var mac_addr string
	dev_rows, dev_err := db.Postgres.Query(context.Background(), `
	SELECT mac_adress, Devices.id FROM Devices 
	INNER JOIN Users ON Users.id = Devices.foreign_id 
	WHERE Users.id = $1;
	`, user)

	for dev_rows.Next() {
		dev_rows.Scan(&mac_addr, &device_id)
		devices_arr = append(devices_arr, Device{
			Id: device_id,
			MacAddr: mac_addr,
		})

		fmt.Printf("Mac: %s\n\n", mac_addr)

	}
	if dev_err != nil {
		middleware.Error(ctx, err, "Erro do db", http.StatusInternalServerError)
		return
	}

	/* -- Encerra READ Devices -- */

	var real_name string
	err = db.Postgres.QueryRow(context.Background(), `
	SELECT real_name FROM users WHERE id = $1 LIMIT 1;
	`, user).Scan(&real_name)

	if err != nil{
		fmt.Printf(err.Error())
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"user": user,
		"contacts": contact_arr,
		"devices": devices_arr,
		"date": date,
	})
}