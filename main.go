package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ShishaDiary struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		testData := []ShishaDiary{
			{Id: 1, Title: "フルーツ系甘め"},
			{Id: 2, Title: "森林系"},
		}
		c.JSON(200, testData)
	})

	router.Run(":8080")
}
