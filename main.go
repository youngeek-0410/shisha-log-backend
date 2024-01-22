package main

import (
	"shisha-log-backend/handler"
	"shisha-log-backend/lib"
	"shisha-log-backend/model/bottle"
	"shisha-log-backend/model/bowl"
	"shisha-log-backend/model/diary"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	diaries := diary.New()
	bottles := bottle.New()
	bowls := bowl.New()
	userBowls := bowl.NewUserBowls()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		os.Getenv("SHISHA_LOG_CLIENT_URL"),
	// 	},
	// }))

	r.GET("/diary", handler.DiariesGet(diaries))
	r.GET("/bottle", handler.BottlesGet(bottles))
	r.GET("/bowl", handler.BowlsGet(bowls))
	r.GET("/user/:user_id/equipment", handler.UserBowlsGet(userBowls))

	r.Run(":8080")
}
