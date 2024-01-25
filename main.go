package main

import (
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	diaries := diary.New()
	userEquipments := handler.NewUserEquipments()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		os.Getenv("SHISHA_LOG_CLIENT_URL"),
	// 	},
	// }))

	r.GET("/diary", handler.DiariesGet(diaries))
	r.GET("/user/:user_id/equipment", handler.UserEquipmentsGet(userEquipments))

	r.Run(":8080")
}
