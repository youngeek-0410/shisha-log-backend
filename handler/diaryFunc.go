package handler

import (
	"net/http"
	"shisha-log-backend/diary"

	"github.com/gin-gonic/gin"
)

func DiariesGet(diaries *diary.Diaries) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := diaries.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

// type DiaryPostRequest struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// }

// func DiaryPost(post *diary.Diaries) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestBody := DiaryPostRequest{}
// 		c.Bind(&requestBody)

// 		item := diary.Diary{
// 			Id:    requestBody.Id,
// 			Title: requestBody.Title,
// 		}
// 		post.Add(item)

// 		c.Status(http.StatusNoContent)
// 	}
// }
