package main

import (
	"log"
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/diary"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	userDiaries := diary.NewUserDiaries()
	userEquipments := handler.NewUserEquipments()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// os.Getenv("SHISHA_LOG_CLIENT_URL"),
			"*",
		},
	}))

	r.GET("/diary/:user_id", handler.GetUserDiaries(userDiaries))
	r.GET("/user/:user_id/equipment", handler.UserEquipmentsGet(userEquipments))
	r.POST("/diary", handler.CreateDiary)

	r.Run(":8080")
}
