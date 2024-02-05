package main

import (
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"
	"shisha-log-backend/model/flavor"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userDiaries := diary.NewUserDiaries()
	userEquipments := handler.NewUserEquipments()

	diaryFlavors := flavor.NewDiaryFlavors()

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

	r.GET("/test/:diary_id", handler.DiaryFlavorsGet(diaryFlavors))

	r.Run(":8080")
}
