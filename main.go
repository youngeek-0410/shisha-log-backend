package main

import (
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userDiaries := diary.NewUserDiaries()
	userEquipments := handler.NewUserEquipments()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		os.Getenv("SHISHA_LOG_CLIENT_URL"),
	// 	},
	// }))

	r.GET("/diary/:user_id", handler.GetUserDiaries(userDiaries))
	r.GET("/user/:user_id/equipment", handler.UserEquipmentsGet(userEquipments))
	r.POST("/diary", handler.CreateDiary)

	r.Run(":8080")
}
