package main

import (
	"log"
	"os"
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

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("SHISHA_LOG_CLIENT_URL"),
			os.Getenv("SHISHA_LOG_CLIENT_STG_URL"),
		},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
	}))

	r.GET("/diary/:user_id", handler.GetUserDiaries(userDiaries))
	r.GET("/user/:user_id/equipment", handler.UserEquipmentsGet(userEquipments))
	r.POST("/diary", handler.CreateDiary)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
