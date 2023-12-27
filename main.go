package main

import (
	"shisha-log-backend/diary"
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	diaries := diary.New()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		os.Getenv("SHISHA_LOG_CLIENT_URL"),
	// 	},
	// }))

	r.GET("/diary", handler.DiariesGet(diaries))

	r.Run(":8080")
}
