package devices

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"quedasegura.com/m/v2/db"
	"quedasegura.com/m/v2/routes/middleware"
)

type NewDevice struct {
	TargetId string `form:"device_id" json:"device_id" xml:"device_id" `
	TargetDevice string `form:"mac_addr" json:"mac_addr" xml:"mac_addr"  binding:"required"`
}

func POST(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, user, _:= middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.Redirect(http.StatusUnauthorized,"/sign_in")
		return
	}

	var device NewDevice

	ctx.ShouldBind(&device)


	db.Postgres.QueryRow(context.Background(), `
	INSERT INTO Devices(id, foreign_id, mac_adress)
	VALUES (
		gen_random_uuid (),
		$1,
		$2
	);
	`, user, device.TargetDevice)
	

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func UPDATE(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, _, _ := middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.Redirect(http.StatusUnauthorized,"/sign_in")
		return
	}

	var device NewDevice

	ctx.ShouldBind(&device)


	db.Postgres.QueryRow(context.Background(), `
	UPDATE devices SET mac_adress = $1 WHERE id = $2;
	`, device.TargetDevice, device.TargetId)
	

	ctx.Redirect(http.StatusMovedPermanently, "/")
}


type DeviceId struct {
	TargetId string `form:"device_id" json:"device_id" xml:"device_id"  binding:"required"`
}

func DELETE(ctx *gin.Context)  {
	cookie, _ := ctx.Cookie("token")

	guard, _, _:= middleware.Guard(ctx, cookie)
		
	if guard == false {
		ctx.Redirect(http.StatusUnauthorized,"/sign_in")
		return
	}

	var device_id DeviceId

	ctx.ShouldBind(&device_id)


	db.Postgres.QueryRow(context.Background(), `
	DELETE FROM devices WHERE id = $1;
	`, device_id.TargetId)

	ctx.Redirect(http.StatusMovedPermanently, "/")
}